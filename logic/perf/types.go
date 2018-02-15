package perf

import "time"

//FROM hey
type Result struct {
	err           error
	statusCode    int
	duration      time.Duration
	connDuration  time.Duration // connection setup(DNS lookup + Dial up) duration
	dnsDuration   time.Duration // dns lookup duration
	reqDuration   time.Duration // request "write" duration
	resDuration   time.Duration // response "read" duration
	delayDuration time.Duration // delay between response and request
	contentLength int64
}

type Report struct {
	avgTotal       float64
	fastest        float64 `json:"fastest"`
	slowest        float64 `json:"slowest"`
	average        float64 `json:"average"`
	rps            float64 `json:"rps"`
	avgConn        float64 `json:"dns_dial"`
	avgDns         float64 `json:"dns_lookup"`
	avgReq         float64 `json:"req_write_avg"`
	avgRes         float64 `json:"res_read_avg"`
	avgDelay       float64 `json:"res_wait_avg"`
	connLats       []float64
	dnsLats        []float64
	reqLats        []float64
	resLats        []float64
	delayLats      []float64
	total          time.Duration
	errorDist      map[string]int `json:"error"`
	statusCodeDist map[int]int
	lats           []float64
	sizeTotal      int64 `json:"size_total"`
}
