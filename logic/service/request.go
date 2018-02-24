package service

import (
	"errors"
	"net/http"
)

type StartPerfRequest struct {
	Number      uint32 `http:"number"`
	Concurrency int    `http:"cc"`
	Qps         int    `http:"qps"`
	Method      string `http:"method"`
	Url         string `http:"url"`
	Body        []byte `http:"body"`
}

type StartUnitRequest struct {
	App    string `http:"app"`
	Router string `http:"router"`
	Cmd    string `http:"cmd"`
}

func checkStartPerfRequest(req StartPerfRequest) bool {
	if req.Number == 0 ||
		req.Concurrency == 0 ||
		req.Qps == 0 ||
		req.Url == "" {
		return false
	}
	return true
}

func ParseStartPerfRequest(req *http.Request) (perfReq StartPerfRequest, err error) {
	err = parsePostRequest(req, &perfReq)
	if err != nil {
		return
	}
	if !checkStartPerfRequest(perfReq) {
		err = errors.New("check start perf query input error")
	}
	return
}

func checkStartUnitRequest(req StartUnitRequest) bool {
	if req.App == "" ||
		req.Router == "" ||
		req.Cmd == "" {
		return false
	}
	return true
}

func ParseStartUnitRequest(req *http.Request) (unitReq StartUnitRequest, err error) {
	parseGetRequest(req, &unitReq, TAG)
	if !checkStartUnitRequest(unitReq) {
		err = errors.New("check start unit query input error")
		return
	}
	return
}
