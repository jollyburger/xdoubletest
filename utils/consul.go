package utils

import (
	"xframe/log"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	TimerTask string `json:"timer_task"`
}

var (
	consulConfig ConsulConfig
	consulClient *api.Client
)

func InitConsul(conf ConsulConfig) {
	consulConfig = conf
	if consulConfig.TimerTask == "" {
		panic("consul configure error")
	}
}

func GetSharedConsulClient() *api.Client {
	if consulClient != nil {
		return consulClient
	}
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.ERRORF("[consul]get consul client error: %v", err)
		return nil
	}
	return consulClient
}
