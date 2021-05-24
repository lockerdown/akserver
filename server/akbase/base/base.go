package base

import "net/http"

// FailReturn api错误返回函数
func FailReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = ""
	res["code"] = http.StatusBadRequest
	res["msg"] = msg

	return res
}

func SuccessReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = msg
	res["code"] = http.StatusOK
	res["msg"] = "success"

	return res
}

func ServerFailReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = ""
	res["code"] = http.StatusInternalServerError
	res["msg"] = msg

	return res
}

func AuthSuccessReturn(data interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = data
	res["code"] = http.StatusOK
	res["msg"] = "认证成功."

	return res
}

func AuthFailReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = ""
	res["code"] = http.StatusUnauthorized
	res["msg"] = msg

	return res
}
