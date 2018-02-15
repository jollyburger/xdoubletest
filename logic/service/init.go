package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r            = mux.NewRouter()
	DEFAULT_PATH = "/test/"
)

var (
	handlerList = map[string]struct {
		Handler func(http.ResponseWriter, *http.Request)
	}{
		"start-perf": {Handler: start_perf},
		"start-unit": {Handler: start_unit},
	}
)

func init() {
	for path, task := range handlerList {
		r.HandleFunc(DEFAULT_PATH+path, task.Handler)
	}
}

func Run(addr string, r_timeout int, w_timeout int) error {
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  r_timeout,
		WriteTimeout: w_timeout,
	}
	err := srv.ListenAndServe()
	return err
}
