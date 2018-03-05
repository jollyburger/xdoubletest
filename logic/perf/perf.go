package perf

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"sort"
	"sync"
	"time"
	"xframe/log"
	"xframe/utils"
)

type Perf interface {
	//start performance task
	Start() (interface{}, error)
	//stop performance task
	Stop() error
}

type DefaultPerf struct {
	Id      string
	Number  uint32
	Cc      int
	Qps     int
	Method  string
	Url     string
	Body    []byte
	stopChs []chan struct{}
	results chan Result
}

func initDefaultPerf(number uint32, cc int, qps int, method string, url string, body []byte) *DefaultPerf {
	this := new(DefaultPerf)
	this.Id = utils.NewUUIDV4().String()
	this.Number = number
	this.Cc = cc
	this.Qps = qps
	this.Method = method
	this.Url = url
	this.Body = body
	this.stopChs = make([]chan struct{}, this.Cc)
	this.results = make(chan Result, this.Number)
	return this
}

func (this *DefaultPerf) makeRequest() (*http.Request, error) {
	if this.Method == "GET" {
		return http.NewRequest(this.Method, this.Url, nil)
	} else if this.Method == "POST" {
		buf := bytes.NewBuffer(this.Body)
		return http.NewRequest(this.Method, this.Url, buf)
	}
	return nil, errors.New("method error")
}

func (this *DefaultPerf) DoRequest(client http.Client) {
	s := time.Now()
	var size int64
	var code int
	var dnsStart, connStart, resStart, reqStart, delayStart time.Time
	var dnsDuration, connDuration, resDuration, reqDuration, delayDuration time.Duration
	req, err := this.makeRequest()
	if err != nil {
		log.ERRORF("make request error: %v", err)
		return
	}
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDuration = time.Now().Sub(dnsStart)
		},
		GetConn: func(h string) {
			connStart = time.Now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			connDuration = time.Now().Sub(connStart)
			reqStart = time.Now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			reqDuration = time.Now().Sub(reqStart)
			delayStart = time.Now()
		},
		GotFirstResponseByte: func() {
			delayDuration = time.Now().Sub(delayStart)
			resStart = time.Now()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := client.Do(req)
	if err == nil {
		size = resp.ContentLength
		code = resp.StatusCode
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	t := time.Now()
	resDuration = t.Sub(resStart)
	finish := t.Sub(s)
	this.results <- Result{
		statusCode:    code,
		duration:      finish,
		err:           err,
		contentLength: size,
		connDuration:  connDuration,
		dnsDuration:   dnsDuration,
		reqDuration:   reqDuration,
		resDuration:   resDuration,
		delayDuration: delayDuration,
	}
}

func (this *DefaultPerf) runWorker(n uint32, stopCh chan struct{}) {
	var counter uint32
	tick := time.NewTicker(time.Duration(1000/this.Qps) * time.Millisecond)
	defer tick.Stop()
	cli := http.Client{}
	for {
		select {
		case <-tick.C:
			counter++
			if counter == n {
				log.DEBUG("counter hit, stop")
				return
			}
			this.DoRequest(cli)
		case <-stopCh:
			log.DEBUG("receive stop signal")
			return
		}
	}
}

func (this *DefaultPerf) Report(t time.Duration) interface{} {
	r := InitReport(t)
	for res := range this.results {
		if res.err != nil {
			r.ErrorDist[res.err.Error()]++
		} else {
			r.lats = append(r.lats, res.duration.Seconds())
			r.avgTotal += res.duration.Seconds()
			r.AvgConn += res.connDuration.Seconds()
			r.AvgDelay += res.delayDuration.Seconds()
			r.AvgDns += res.dnsDuration.Seconds()
			r.AvgReq += res.reqDuration.Seconds()
			r.AvgRes += res.resDuration.Seconds()
			r.connLats = append(r.connLats, res.connDuration.Seconds())
			r.dnsLats = append(r.dnsLats, res.dnsDuration.Seconds())
			r.reqLats = append(r.reqLats, res.reqDuration.Seconds())
			r.delayLats = append(r.delayLats, res.delayDuration.Seconds())
			r.resLats = append(r.resLats, res.resDuration.Seconds())
			r.statusCodeDist[res.statusCode]++
			if res.contentLength > 0 {
				r.SizeTotal += res.contentLength
			}
		}
	}
	r.Rps = float64(len(r.lats)) / r.total.Seconds()
	r.Average = r.avgTotal / float64(len(r.lats))
	r.AvgConn = r.AvgConn / float64(len(r.lats))
	r.AvgDelay = r.AvgDelay / float64(len(r.lats))
	r.AvgDns = r.AvgDns / float64(len(r.lats))
	r.AvgReq = r.AvgReq / float64(len(r.lats))
	r.AvgRes = r.AvgRes / float64(len(r.lats))
	if len(r.lats) != 0 {
		sort.Float64s(r.lats)
		r.Fastest = r.lats[0]
		r.Slowest = r.lats[len(r.lats)-1]
	}
	return r
}

func (this *DefaultPerf) Start() (interface{}, error) {
	//split into cc worker with number / cc request
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(this.Cc)
	for i := 0; i < this.Cc; i++ {
		go func(i int) {
			this.runWorker(this.Number/uint32(this.Cc), this.stopChs[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(this.results)
	return this.Report(time.Now().Sub(start)), nil
}

func (this *DefaultPerf) Stop() error {
	for _, ch := range this.stopChs {
		if ch != nil {
			go close(ch)
		}
	}
	return nil
}
