package service

const (
	SUCCESS     = iota
	INPUT_ERROR = iota + 6000
)

var (
	errorMap = map[int]string{
		INPUT_ERROR: "request query error",
	}
)
