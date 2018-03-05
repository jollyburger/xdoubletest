package perf

import (
	"sync"
	"xframe/log"
)

var (
	GperfManager = InitPerfManager()
)

type PerfManager struct {
	sync.RWMutex
	Tasks map[string]*DefaultPerf
}

func InitPerfManager() *PerfManager {
	this := new(PerfManager)
	this.Tasks = make(map[string]*DefaultPerf)
	return this
}

func (this *PerfManager) add(perf *DefaultPerf) {
	this.Lock()
	this.Tasks[perf.Id] = perf
	this.Unlock()
}

func (this *PerfManager) remove(tid string) {
	this.Lock()
	if _, ok := this.Tasks[tid]; ok {
		//t.Stop()
		delete(this.Tasks, tid)
	}
	this.Unlock()
}

func (this *PerfManager) Start(number uint32, cc int, qps int, method string, url string, body []byte) (interface{}, error) {
	dp := initDefaultPerf(number, cc, qps, method, url, body)
	this.add(dp)
	data, err := dp.Start()
	this.remove(dp.Id)
	return data, err
}

func (this *PerfManager) Stop(tid string) error {
	this.remove(tid)
	return nil
}

func (this *PerfManager) TasksLength() int {
	this.Lock()
	defer this.Unlock()
	return len(this.Tasks)
}

func (this *PerfManager) DumpTasks() {
	this.RLock()
	for id, perf := range this.Tasks {
		log.DEBUGF("Task %d, Id: %s, Request Num: %d, Concurrency: %d, Qps: %d", id, perf.Id, perf.Number, perf.Cc, perf.Qps)
	}
	this.RUnlock()
}
