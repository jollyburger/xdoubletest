package utils

import (
	"fmt"
	"xframe/server"
)

type HttpConfig struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

func (this HttpConfig) Addr() (addr string) {
	var ip string
	ip, err := server.ParseListenAddr(this.Address)
	if err != nil {
		return
	}
	addr = fmt.Sprintf("%s:%d", ip, this.Port)
	return
}

var (
	httpConfig HttpConfig
)

func InitHttp(conf HttpConfig) HttpConfig {
	httpConfig = conf
	if httpConfig.Address == "" ||
		httpConfig.Port == 0 ||
		httpConfig.ReadTimeout == 0 ||
		httpConfig.WriteTimeout == 0 {
		panic("http server configure error")
	}
	return httpConfig
}
