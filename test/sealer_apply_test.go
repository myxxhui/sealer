package test

import (
	"fmt"

	"github.com/alibaba/sealer/test/suites/apply"
	"github.com/alibaba/sealer/test/suites/registry"
	"github.com/alibaba/sealer/test/testhelper"
	"github.com/alibaba/sealer/test/testhelper/settings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("sealer apply", func() {
	Context("start apply", func() {
		BeforeEach(func() {
			registry.Login()
		})

		Context("with roofs images", func() {
			clusterFile := apply.GetClusterFilePathOfRootfs()
			AfterEach(func() {
				cluster := apply.GetClusterFileData(clusterFile)
				apply.DeleteCluster(cluster.ClusterName)
			})

			It("apply cluster", func() {
				sess, err := testhelper.Start(fmt.Sprintf("sealer apply -f %s", clusterFile))
				Expect(err).NotTo(HaveOccurred())
				Eventually(sess, settings.MaxWaiteTime).Should(Exit(0))
			})
		})

	})

})
