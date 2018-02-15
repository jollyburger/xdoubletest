package service

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

var (
	TAG = "http"
)

func convertType(sf reflect.StructField, rv reflect.Value, value string) {
	var v interface{}
	switch sf.Type.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32:
		tmp_v, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		v = tmp_v
	case reflect.Int64:
		tmp_v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return
		}
		v = tmp_v
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tmp_v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return
		}
		v = tmp_v
	case reflect.Float32, reflect.Float64:
		tmp_v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		v = tmp_v
	case reflect.String:
		v = value
	case reflect.Bool:
		v = (value == "1")
	}
	rv.Set(reflect.ValueOf(v))
}

func parseGetRequest(req *http.Request, obj interface{}, tag string) {
	v := reflect.Indirect(reflect.ValueOf(obj))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		rv := v.Field(i)
		if _, ok := sf.Tag.Lookup(tag); !ok {
			continue
		}
		st := sf.Tag.Get(tag)
		reqV := req.FormValue(st)
		if reqV != "" {
			convertType(sf, rv, reqV)
		}
	}
}

func parsePostRequest(req *http.Request, obj interface{}) (err error) {
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&obj)
	if err != nil {
		return
	}
	defer req.Body.Close()
	return
}
