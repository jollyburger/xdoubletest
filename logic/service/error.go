package service

const (
	SUCCESS     = iota
	INPUT_ERROR = iota + 6000
	PERF_TASK_ERROR
	UNIT_TASK_ERROR
)

var (
	ERROR_MAP = map[int]string{
		INPUT_ERROR:     "request query error",
		PERF_TASK_ERROR: "performance test error",
		UNIT_TASK_ERROR: "unit test error",
	}
)
