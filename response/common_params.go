package response

type CommonParams struct {
	RequestId string `举例：	8906582E-6722	请求ID`
	Code      string `举例：	OK	状态码-返回OK代表请求成功,其他错误码详见错误码列表`
	Message   string `举例：	请求成功	状态码的描述`
}
