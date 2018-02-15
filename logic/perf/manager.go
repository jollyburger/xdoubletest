package perf

import "sync"

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
}

func (this *PerfManager) add(perf *DefaultPerf) {
	this.Lock()
	this.Tasks[perf.Id] = perf
	this.Unlock()
}

func (this *PerfManager) remove(tid string) {
	this.Lock()
	if _, ok := this.Tasks[tid]; ok {
		perf.Stop()
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
