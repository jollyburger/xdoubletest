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
	avgTotal       float64        `json:"-"`
	Fastest        float64        `json:"fastest"`
	Slowest        float64        `json:"slowest"`
	Average        float64        `json:"average"`
	Rps            float64        `json:"rps"`
	AvgConn        float64        `json:"dns_dial"`
	AvgDns         float64        `json:"dns_lookup"`
	AvgReq         float64        `json:"req_write_avg"`
	AvgRes         float64        `json:"res_read_avg"`
	AvgDelay       float64        `json:"res_wait_avg"`
	connLats       []float64      `json:"-"`
	dnsLats        []float64      `json:"-"`
	reqLats        []float64      `json:"-"`
	resLats        []float64      `json:"-"`
	delayLats      []float64      `json:"-"`
	total          time.Duration  `json:"-"`
	ErrorDist      map[string]int `json:"error"`
	statusCodeDist map[int]int    `json:"-"`
	lats           []float64      `json:"-"`
	SizeTotal      int64          `json:"size_total"`
}

func InitReport(t time.Duration) Report {
	var r Report
	r.total = t
	r.ErrorDist = make(map[string]int)
	r.statusCodeDist = make(map[int]int)
	r.connLats = make([]float64, 0)
	r.dnsLats = make([]float64, 0)
	r.reqLats = make([]float64, 0)
	r.resLats = make([]float64, 0)
	r.delayLats = make([]float64, 0)
	r.lats = make([]float64, 0)
	return r
}
