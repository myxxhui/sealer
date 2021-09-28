package runtime

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/alibaba/sealer/common"

	"github.com/alibaba/sealer/logger"

	"github.com/alibaba/sealer/cert"
	"github.com/alibaba/sealer/guest"
	v1 "github.com/alibaba/sealer/types/api/v1"
	"github.com/alibaba/sealer/utils"
	"github.com/alibaba/sealer/utils/ssh"
)

const (
	RemoteCmdInitEtcdDir   = "mkdir -p /var/lib/etcd && mount %s /var/lib/etcd && rm -rf /var/lib/etcd/* && echo \"%s /var/lib/etcd ext4 defaults 0 0\" >> /etc/fstab"
	RemoteCmdUnmountEtcd   = "umount /var/lib/etcd; mkfs.ext4 -F %s"
	RemoteCmdLinkStatic    = "mkdir -p %s && ln -f %s %s"
	RemoteApplyYaml        = `echo '%s' | kubectl apply -f -`
	WriteKubeadmConfigCmd  = "cd %s && echo \"%s\" > kubeadm-config.yaml"
	DefaultVIP             = "10.103.97.2"
	DefaultAPIserverDomain = "apiserver.cluster.local"
	DefaultRegistryPort    = 5000
)

func (d *Default) init(cluster *v1.Cluster) error {
	/*
		if err := d.initRunner(cluster); err != nil {
			return err
		}
	*/
	//config kubeadm
	if err := d.ConfigKubeadmOnMaster0(); err != nil {
		return err
	}

	//generate certs
	if err := d.GenerateCert(); err != nil {
		return err
	}

	//create kubeConfig for master0
	if err := d.CreateKubeConfig(); err != nil {
		return err
	}

	if err := d.LinkStaticFiles(d.Masters); err != nil {
		return err
	}

	if err := d.EnsureRegistryOnMaster0(); err != nil {
		return err
	}

	if err := d.InitMaster0(); err != nil {
		return err
	}

	return nil
}

func (d *Default) initRunner(cluster *v1.Cluster) error {
	d.SSH = ssh.NewSSHByCluster(cluster)
	d.ClusterName = cluster.Name
	d.SvcCIDR = cluster.Spec.Network.SvcCIDR
	d.PodCIDR = cluster.Spec.Network.PodCIDR
	// TODO add host port
	d.Masters = cluster.Spec.Masters.IPList
	d.VIP = DefaultVIP
	d.RegistryPort = DefaultRegistryPort
	// TODO add host port
	d.Nodes = cluster.Spec.Nodes.IPList
	d.APIServer = DefaultAPIserverDomain
	d.Rootfs = path.Join(common.DefaultClusterRootfsDir, d.ClusterName)
	d.CertPath = fmt.Sprintf("%s/pki", d.Rootfs)
	d.CertEtcdPath = fmt.Sprintf("%s/etcd", d.CertPath)
	d.StaticFileDir = fmt.Sprintf("%s/statics", d.Rootfs)
	// TODO remote port in ipList
	d.APIServerCertSANs = append(cluster.Spec.CertSANS, d.getDefaultSANs()...)
	d.Interface = cluster.Spec.Network.Interface
	d.Network = cluster.Spec.Network.CNIName
	d.PodCIDR = cluster.Spec.Network.PodCIDR
	d.SvcCIDR = cluster.Spec.Network.SvcCIDR
	d.WithoutCNI = cluster.Spec.Network.WithoutCNI
	if d.IPIP && d.MTU == "" {
		d.MTU = "1480"
	} else {
		d.MTU = "1550"
	}

	return nil
}
func (d *Default) ConfigKubeadmOnMaster0() error {
	var templateData string
	var err error
	var tpl []byte
	var fileData []byte
	if d.KubeadmFilePath == "" {
		tpl, err = d.defaultTemplate()
	} else {
		//TODO rootfs kubeadm.tmpl
		fileData, err = ioutil.ReadFile(d.KubeadmFilePath)
		if err != nil {
			return err
		}
		tpl, err = d.templateFromContent(string(fileData))
	}

	if err != nil {
		return err
	}
	templateData = string(tpl)

	cmd := fmt.Sprintf(WriteKubeadmConfigCmd, d.Rootfs, templateData)
	err = d.SSH.CmdAsync(d.Masters[0], cmd)
	if err != nil {
		return err
	}

	kubeadm := kubeadmDataFromYaml(templateData)
	if kubeadm != nil {
		d.DNSDomain = kubeadm.Networking.DNSDomain
		d.APIServerCertSANs = kubeadm.APIServer.CertSANs
	} else {
		logger.Warn("decode certSANs from config failed, using default SANs")
		d.APIServerCertSANs = d.getDefaultSANs()
	}
	return nil
}

func (d *Default) GenerateCert() error {
	err := cert.GenerateCert(
		d.CertPath,
		d.CertEtcdPath,
		d.APIServerCertSANs,
		utils.GetHostIP(d.Masters[0]),
		d.GetRemoteHostName(d.Masters[0]),
		d.SvcCIDR,
		d.DNSDomain,
	)
	if err != nil {
		return fmt.Errorf("generate certs failed %v", err)
	}
	d.sendNewCertAndKey(d.Masters[:1])

	return nil
}

func (d *Default) CreateKubeConfig() error {
	hostname := d.GetRemoteHostName(d.Masters[0])

	certConfig := cert.Config{
		Path:     d.CertPath,
		BaseName: "ca",
	}

	controlPlaneEndpoint := fmt.Sprintf("https://%s:6443", d.APIServer)

	err := cert.CreateJoinControlPlaneKubeConfigFiles(d.Rootfs,
		certConfig, hostname, controlPlaneEndpoint, "kubernetes")
	if err != nil {
		return fmt.Errorf("generator kubeconfig failed %s", err)
	}
	return nil
}

//InitMaster0 is
func (d *Default) InitMaster0() error {
	d.SendJoinMasterKubeConfigs(d.Masters[:1], AdminConf, ControllerConf, SchedulerConf, KubeletConf)
	/*
		err := d.mountEtcdDisk(d.Masters[:1], d.EtcdDevice)
		if err != nil {
			return fmt.Error("mount for /var/lib/etcd failed at %s, due to %s", d.Masters[0], err)
		}
	*/

	cmdAddEtcHost := fmt.Sprintf(RemoteAddEtcHosts, getAPIServerHost(utils.GetHostIP(d.Masters[0]), d.APIServer))
	cmdAddRegistryHosts := fmt.Sprintf(RemoteAddEtcHosts, getRegistryHost(utils.GetHostIP(d.Masters[0])))
	err := d.SSH.CmdAsync(d.Masters[0], cmdAddEtcHost, cmdAddRegistryHosts)
	if err != nil {
		return err
	}

	cmdInit := d.Command(d.Metadata.Version, InitMaster)

	// TODO skip docker version error check for test
	output, err := d.SSH.Cmd(d.Masters[0], fmt.Sprintf("%s --ignore-preflight-errors=SystemVerification", cmdInit))
	if err != nil {
		return fmt.Errorf("init master0 failed, error: %s. Please clean and reinstall", err.Error())
	}
	d.decodeMaster0Output(output)
	err = d.SSH.CmdAsync(d.Masters[0], RemoteCopyKubeConfig)
	if err != nil {
		return err
	}

	return d.InitCNI()
}

func (d *Default) InitCNI() error {
	logger.Info("start to install CNI")
	if d.WithoutCNI {
		return nil
	}
	// can-reach is used by calico multi network , flannel has nothing to add. just Use it.
	if len(strings.Split(d.Interface, ".")) == 4 && d.Network == "calico" {
		d.Interface = "can-reach=" + d.Interface
	} else {
		d.Interface = "interface=" + d.Interface
	}

	netYaml, err := guest.NewNetWork(d.Network, guest.MetaData{
		Interface: d.Interface,
		CIDR:      d.PodCIDR,
		IPIP:      d.IPIP,
		MTU:       d.MTU,
	}).Manifests("")
	if err != nil {
		return fmt.Errorf("render net manifests failed %v", err)
	}
	logger.Info("render cni yaml success")

	return d.SSH.CmdAsync(d.Masters[0], fmt.Sprintf(RemoteApplyYaml, netYaml))
}

/*func (d *Default) mountEtcdDisk(targetHosts []string, etcdDisk string) error {
	if etcdDisk == "" {
		logger.Warn("Etcd Disk is not set, etcd now uses root disk which is not recommended due to stability requirement.")
		return nil
	}

	var wg sync.WaitGroup
	for _, host := range targetHosts {
		wg.Add(1)
		go func(master string) {
			defer wg.Done()
			cmdInitDevice := fmt.Sprintf(RemoteCmdUnmountEtcd, etcdDisk)
			cmdInitDir := fmt.Sprintf(RemoteCmdInitEtcdDir, etcdDisk, etcdDisk)
			err := d.SSH.CmdAsync(master, cmdInitDevice, cmdInitDir)
			if err != nil {
				logger.Error("[%s] mount %s /var/lib/etcd failed, please check disk configuration", master, etcdDisk)
				os.Exit(1)
			}
		}(host)
	}
	wg.Wait()

	return nil
}*/

func (d *Default) LinkStaticFiles(nodes []string) error {
	var flag bool
	for _, file := range MasterStaticFiles {
		staticFilePath := filepath.Join(d.StaticFileDir, file.Name)
		cmdLinkStatic := fmt.Sprintf(RemoteCmdLinkStatic, file.DestinationDir, staticFilePath, filepath.Join(file.DestinationDir, file.Name))
		var wg sync.WaitGroup
		for _, host := range nodes {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				err := d.SSH.CmdAsync(host, cmdLinkStatic)
				if err != nil {
					logger.Error("[%s] link static file failed, error:%s", host, err.Error())
					flag = true
				}
			}(host)
			if flag {
				return fmt.Errorf("link static files failed %s %s", host, cmdLinkStatic)
			}
		}
		wg.Wait()
	}
	return nil
}

//decode output to join token  hash and key
func (d *Default) decodeMaster0Output(output []byte) {
	s0 := string(output)
	logger.Debug("[globals]decodeOutput: %s", s0)
	slice := strings.Split(s0, "kubeadm join")
	slice1 := strings.Split(slice[1], "Please note")
	logger.Info("[globals]join command is: %s", slice1[0])
	d.decodeJoinCmd(slice1[0])
}

//  192.168.0.200:6443 --token 9vr73a.a8uxyaju799qwdjv --discovery-token-ca-cert-hash sha256:7c2e69131a36ae2a042a339b33381c6d0d43887e2de83720eff5359e26aec866 --experimental-control-plane --certificate-key f8902e114ef118304e561c3ecd4d0b543adc226b7a07f675f56564185ffe0c07
func (d *Default) decodeJoinCmd(cmd string) {
	logger.Debug("[globals]decodeJoinCmd: %s", cmd)
	stringSlice := strings.Split(cmd, " ")

	for i, r := range stringSlice {
		// upstream error, delete \t, \\, \n, space.
		r = strings.ReplaceAll(r, "\t", "")
		r = strings.ReplaceAll(r, "\n", "")
		r = strings.ReplaceAll(r, "\\", "")
		r = strings.TrimSpace(r)
		if strings.Contains(r, "--token") {
			d.JoinToken = stringSlice[i+1]
		}
		if strings.Contains(r, "--discovery-token-ca-cert-hash") {
			d.TokenCaCertHash = stringSlice[i+1]
		}
		if strings.Contains(r, "--certificate-key") {
			d.CertificateKey = stringSlice[i+1][:64]
		}
	}
	logger.Debug("joinToken: %v\nTokenCaCertHash: %v\nCertificateKey: %v", d.JoinToken, d.TokenCaCertHash, d.CertificateKey)
}
