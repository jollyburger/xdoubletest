package service

import (
	"encoding/json"
	"net/http"
	"xframe/log"
)

type BaseResponse struct {
	RetCode int    `json:"retcode"`
	Message string `json:"message"`
}

type DataResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

func DoBaseResponse(rw http.ResponseWriter, retCode int) {
	var res BaseResponse
	res.RetCode = -1 * retCode
	if msg, ok := ERROR_MAP[retCode]; ok {
		res.Message = msg
	}
	buf, err := json.Marshal(res)
	if err != nil {
		log.ERROR("marshal json response error:", err)
		return
	}
	rw.Write(buf)
}

func DoDataResponse(rw http.ResponseWriter, retCode int, data interface{}) {
	var res DataResponse
	res.RetCode = -1 * retCode
	if msg, ok := ERROR_MAP[retCode]; ok {
		res.Message = msg
	}
	res.Data = data
	buf, err := json.Marshal(res)
	if err != nil {
		log.ERROR("marshal json response error:", err)
		return
	}
	rw.Write(buf)
}
