package infra

import (
	"github.com/alibaba/sealer/infra/aliyun"
	"github.com/alibaba/sealer/logger"
	v1 "github.com/alibaba/sealer/types/api/v1"
)

type Interface interface {
	// Apply apply iaas resources and save metadata info like vpc instance id to cluster status
	// https://github.com/fanux/sealgate/tree/master/cloud
	Apply() error
}

func NewDefaultProvider(cluster *v1.Cluster) Interface {
	switch cluster.Spec.Provider {
	case aliyun.AliCloud:
		config := new(aliyun.Config)
		err := aliyun.LoadConfig(config)
		if err != nil {
			logger.Error(err)
			return nil
		}
		aliProvider := new(aliyun.AliProvider)
		aliProvider.Config = *config
		aliProvider.Cluster = cluster
		err = aliProvider.NewClient()
		if err != nil {
			logger.Error(err)
		}
		return aliProvider
	default:
		return nil
	}
}
