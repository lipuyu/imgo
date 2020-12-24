package re

import "stouch_server/common/er"

type ResponseResult struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func NewByData(data interface{}) ResponseResult {
	return ResponseResult{Status: true, Msg:"", Data:data, Code: 0}
}

func NewByError(error er.Error) ResponseResult {
	result := ResponseResult{Status: false, Msg: error.Msg, Code: error.Code}
	return result
}
