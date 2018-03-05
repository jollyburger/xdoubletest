package main

import (
	"flag"
	"fmt"
	"xdoubletest/app"
	"xdoubletest/logic/service"
	"xdoubletest/utils"
	"xframe/metric"
)

var (
	conf = flag.String("c", "", "configuration file path")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	var config app.Config
	err := app.ReadConf(&config, *conf)
	if err != nil {
		panic("read conf error")
	}
	fmt.Println(config)
	//init config
	utils.InitLog(config.LogConf)
	utils.InitConsul(config.ConsulConf)
	httpConf := utils.InitHttp(config.HttpConf)
	/*
	 * temporary pprof
	 */
	go func() {
		metric.InitPprof("0.0.0.0:6001")
	}()
	//start service
	if err := service.Run(httpConf.Addr(), httpConf.ReadTimeout, httpConf.WriteTimeout); err != nil {
		panic(err)
	}
}
