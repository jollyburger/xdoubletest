package service

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	r            = mux.NewRouter()
	DEFAULT_PATH = "/test/"
)

var (
	handlerList = map[string]struct {
		Handler func(*http.Request) (interface{}, int)
	}{
		"start-perf": {Handler: start_perf},
		"start-unit": {Handler: start_unit},
	}
)

func httpWrapper(f func(*http.Request) (interface{}, int)) func(http.ResponseWriter, *http.Request) {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		data, err := f(r)
		if err != SUCCESS {
			DoBaseResponse(rw, err)
		} else {
			if data != nil {
				DoDataResponse(rw, SUCCESS, data)
			} else {
				DoBaseResponse(rw, SUCCESS)
			}
		}
	}
	return fn
}

func init() {
	for path, task := range handlerList {
		r.HandleFunc(DEFAULT_PATH+path, httpWrapper(task.Handler))
	}
}

func Run(addr string, r_timeout int, w_timeout int) error {
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  time.Duration(r_timeout) * time.Second,
		WriteTimeout: time.Duration(w_timeout) * time.Second,
	}
	err := srv.ListenAndServe()
	return err
}
