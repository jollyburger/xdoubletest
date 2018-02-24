package unit

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"xdoubletest/utils"
	"xframe/server"
)

var (
	GunitManager = InitUnit()
)

type Unit struct {
	Cpu     int
	CpuPer  int
	Mem     int
	MemSwap int
	Timeout int
}

func InitUnit() *Unit {
	return &Unit{
		Cpu:     DEFAULT_CPU,
		CpuPer:  DEFAULT_CPU_PERCENT,
		Mem:     DEFAULT_MEM,
		MemSwap: DEFAULT_MEM_SWAP,
		Timeout: DEFAULT_TIMEOUT,
	}
}

func (this *Unit) RegisterTestSpace(appName string, desc string) error {
	return nil
}

func (this *Unit) ActivateTestSpace(appName string) error {
	return nil
}

func (this *Unit) RegisterTestCase(appName string, routerName string, method string, imageName string, tag string, desc string) error {
	return nil
}

func (this *Unit) Start(app string, router string, cmd string) (interface{}, error) {
	addr := fmt.Sprintf("http://%s/%s/%s", utils.GetVfAddr(), app, router)
	params := map[string]interface{}{
		"Cpu":        this.Cpu,
		"CpuPercent": this.CpuPer,
		"Memory":     this.Mem,
		"MemorySwap": this.MemSwap,
		"Timeout":    this.Timeout,
		"Cmd":        cmd,
	}
	result, err := server.SendHTTPRequest(nil, addr, params, uint32(this.Timeout))
	if err != nil {
		return nil, err
	}
	var vfRes VfResponse
	err = json.Unmarshal(result, &vfRes)
	if err != nil {
		return nil, err
	}
	if vfRes.RetCode != 0 {
		return nil, errors.New(vfRes.ErrMessage)
	}
	if vfRes.Result.Stdout != "" {
		data, err := base64.StdEncoding.DecodeString(vfRes.Result.Stdout)
		return data, err
	} else if vfRes.Result.Stderr != "" {
		data, err := base64.StdEncoding.DecodeString(vfRes.Result.Stderr)
		return data, err
	}
	return nil, errors.New("output is empty")
}
