package ssh

import (
	"testing"

	"github.com/alibaba/sealer/logger"
)

func TestSSHCopyLocalToRemote(t *testing.T) {
	type args struct {
		host       string
		localPath  string
		remotePath string
	}
	var (
		host = "10.96.33.168"
		ssh  = SSH{
			User:       "root",
			Password:   "123456",
			PkFile:     "",
			PkPassword: "",
			Timeout:    nil,
		}
	)
	tests := []struct {
		name    string
		fields  SSH
		args    args
		wantErr bool
	}{
		{
			name:   "test for copy file to remote server",
			fields: ssh,
			args: args{
				host,
				"../test/file/01",
				"/home/temp/01",
			},
			wantErr: false,
		},
		{
			name:   "test copy for file when local file is not exist",
			fields: ssh,
			args: args{
				host,
				// local file  001 is not exist.
				"../test/file/123",
				"/home/temp/01",
			},
			wantErr: false,
		},
		{
			name:   "test copy dir to remote server",
			fields: ssh,
			args: args{
				host,
				"../test/file",
				"/home/temp011",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SSH{
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				PkFile:     tt.fields.PkFile,
				PkPassword: tt.fields.PkPassword,
				Timeout:    tt.fields.Timeout,
			}

			if !fileExist(tt.args.localPath) {
				logger.Error("local filepath is not exit")
				return
			}
			//if ss.IsFileExist(host, tt.args.remotePath) {
			//	log.Error("remote filepath is exit")
			//	return
			//}
			// test copy dir
			err := ss.Copy(tt.args.host, tt.args.localPath, tt.args.remotePath)
			if (err != nil) != tt.wantErr {
				logger.Error(err)
				t.Errorf("err: %v", err)
			}

			// test the copy result
			//ss.Cmd(tt.args.host, "ls -lh "+tt.args.remotePath)

			// rm remote file
			//ss.Cmd(tt.args.host, "rm -rf "+tt.args.remotePath)
		})
	}
}

func TestSSHFetchRemoteToLocal(t *testing.T) {
	type args struct {
		host       string
		localPath  string
		remotePath string
	}
	var (
		host = ""
		ssh  = SSH{
			User:       "root",
			Password:   "",
			PkFile:     "",
			PkPassword: "",
			Timeout:    nil,
		}
	)
	tests := []struct {
		name    string
		fields  SSH
		args    args
		wantErr bool
	}{
		{
			name:   "test for fetch remote file to local",
			fields: ssh,
			args: args{
				host,
				"/root/.kube/config",
				"/root/Clusterfile",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SSH{
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				PkFile:     tt.fields.PkFile,
				PkPassword: tt.fields.PkPassword,
				Timeout:    tt.fields.Timeout,
			}

			if !ss.IsFileExist(host, tt.args.remotePath) {
				logger.Error("remote filepath is not exit")
				return
			}
			err := ss.Fetch(tt.args.host, tt.args.localPath, tt.args.remotePath)
			if (err != nil) != tt.wantErr {
				logger.Error(err)
				t.Errorf("err: %v", err)
			}
		})
	}
}

func TestSSH_Copy(t *testing.T) {
	type fields struct {
		User       string
		Password   string
		PkFile     string
		PkPassword string
	}
	type args struct {
		host       string
		localPath  string
		remotePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test copy dir",
			fields{
				User:       "root",
				Password:   "",
				PkFile:     "",
				PkPassword: "",
			},
			args{
				host:       "",
				localPath:  "./pkg/cert/pki",
				remotePath: "/root/kubernetes/pki",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SSH{
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				PkFile:     tt.fields.PkFile,
				PkPassword: tt.fields.PkPassword,
			}
			if err := s.Copy(tt.args.host, tt.args.localPath, tt.args.remotePath); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
