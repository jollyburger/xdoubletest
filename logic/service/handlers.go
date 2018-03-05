package service

import (
	"net/http"
	"xdoubletest/logic/perf"
	"xdoubletest/logic/unit"
	"xframe/log"
)

func start_perf(r *http.Request) (interface{}, int) {
	req, err := ParseStartPerfRequest(r)
	if err != nil {
		log.ERRORF("[start_perf]parse input error: %v", err)
		return nil, INPUT_ERROR
	}
	//start perf
	data, err := perf.GperfManager.Start(req.Number, req.Concurrency, req.Qps, req.Method, req.Url, req.Body)
	if err != nil {
		log.ERRORF("[start_perf]perf task error: %v", err)
		return nil, PERF_TASK_ERROR
	}
	log.DEBUGF("[start_perf]perf manager task list length: %d", perf.GperfManager.TasksLength())
	perf.GperfManager.DumpTasks()
	log.INFOF("[start_perf]complete start perf request: %v", data)
	return data, SUCCESS
}

func start_unit(r *http.Request) (interface{}, int) {
	req, err := ParseStartUnitRequest(r)
	if err != nil {
		log.ERRORF("[start_unit]parse input error: %v", err)
		return nil, INPUT_ERROR
	}
	//start unit
	data, err := unit.GunitManager.Start(req.App, req.Router, req.Cmd)
	if err != nil {
		log.ERRORF("[start_unit]unit task error: %v", err)
		return nil, UNIT_TASK_ERROR
	}
	log.INFOF("[start_unit]complete start unit request")
	return data, SUCCESS
}
