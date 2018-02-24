package main

import (
	"flag"
	"xdoubletest/app"
	"xdoubletest/logic/service"
	"xdoubletest/utils"
)

var (
	conf = flag.String("c", "examples/test.json", "configuration file path")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	var config app.Config
	err := app.ReadConf(config, *conf)
	if err != nil {
		panic("read conf error")
	}
	//init config
	utils.InitLog(config.LogConf)
	utils.InitConsul(config.ConsulConf)
	httpConf := utils.InitHttp(config.HttpConf)
	//start service
	if err := service.Run(httpConf.Addr(), httpConf.ReadTimeout, httpConf.WriteTimeout); err != nil {
		panic(err)
	}
}
