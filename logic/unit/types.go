package unit

var (
	DEFAULT_CPU         = 1
	DEFAULT_CPU_PERCENT = 50
	DEFAULT_MEM         = 100
	DEFAULT_MEM_SWAP    = 100
	DEFAULT_TIMEOUT     = 10
)

type Output struct {
	Stdout string `json:"StdOut"`
	Stderr string `json:"StdErr"`
	File   []byte `json:"File"`
}

type VfResponse struct {
	Action     string `json:"Action"`
	RetCode    int    `json:"RetCode"`
	ErrMessage string `json:"ErrMessage"`
	Result     Output `json:"Result"`
}
