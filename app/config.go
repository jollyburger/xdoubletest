package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"xdoubletest/utils"
)

type Config struct {
	HttpConf   utils.HttpConfig   `json:"http"`
	LogConf    utils.LogConfig    `json:"log"`
	ConsulConf utils.ConsulConfig `json:"consul"`
	//PprofConf  utils.PprofConfig  `json:"pprof"`
	VfConf utils.VfConfig `json:"vf"`
}

func ReadConf(config Config, filePath string) (err error) {
	confBuf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("load config file %s error: %v\n", filePath, err)
		return
	}
	err = json.Unmarshal(confBuf, &config)
	if err != nil {
		fmt.Printf("parse config file %s error: %v\n", filePath, err)
		return
	}
	return
}
