package result

import "encoding/json"

type Results struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// Success 成功返回结果
func Success(data interface{}) Results {
	return Results{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

// Fail 失败返回结果
func Fail(msg string) Results {
	return Results{
		Code: 400,
		Msg:  msg,
		Data: nil,
	}
}

// FailWithCode 自定义失败返回结果
func FailWithCode(code int, msg string) Results {
	return Results{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// FailWithDetailed 失败返回结果，带有详细信息
func FailWithDetailed(code int, msg string, data interface{}) Results {
	return Results{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Results) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func result(code int, msg string) *Results {
	return &Results{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
