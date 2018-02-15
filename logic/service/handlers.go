package service

import (
	"net/http"
	"xframe/log"
)

func start_perf(rw http.ResponseWriter, r *http.Request) {
	req, err := ParseStartPerfRequest(r)
	if err != nil {
		log.ERRORF("[start_perf]parse input error: %v", err)
		DoBaseResponse(rw, INPUT_ERROR)
		return
	}
	//start perf
	log.INFOF("[start_perf]complete start perf request")
	DoDataResponse(rw, SUCCESS, data)
	return
}

func start_unit(rw http.ResponseWriter, r *http.Request) {
	req, err := ParseStartUnitRequest(r)
	if err != nil {
		log.ERRORF("[start_unit]parse input error: %v", err)
		DoBaseResponse(rw, INPUT_ERROR)
		return
	}
	//start unit
	log.INFOF("[start_unit]complete start unit request")
	DoDataResponse(rw, SUCCESS, data)
	return
}
