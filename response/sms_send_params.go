package response

/**
https://help.aliyun.com/document_detail/55284.html

Code	描述
OK	请求成功
isp.RAM_PERMISSION_DENY	RAM权限DENY
isv.OUT_OF_SERVICE	业务停机
isv.PRODUCT_UN_SUBSCRIPT	未开通云通信产品的阿里云客户
isv.PRODUCT_UNSUBSCRIBE	产品未开通
isv.ACCOUNT_NOT_EXISTS	账户不存在
isv.ACCOUNT_ABNORMAL	账户异常
isv.SMS_TEMPLATE_ILLEGAL	短信模板不合法
isv.SMS_SIGNATURE_ILLEGAL	短信签名不合法
isv.INVALID_PARAMETERS	参数异常
isp.SYSTEM_ERROR	系统错误
isv.MOBILE_NUMBER_ILLEGAL	非法手机号
isv.MOBILE_COUNT_OVER_LIMIT	手机号码数量超过限制
isv.TEMPLATE_MISSING_PARAMETERS	模板缺少变量
isv.BUSINESS_LIMIT_CONTROL	业务限流
isv.INVALID_JSON_PARAM	JSON参数不合法，只接受字符串值
isv.BLACK_KEY_CONTROL_LIMIT	黑名单管控
isv.PARAM_LENGTH_LIMIT	参数超出长度限制
isv.PARAM_NOT_SUPPORT_URL	不支持URL
isv.AMOUNT_NOT_ENOUGH	账户余额不足
*/

type SmsSendParams struct {
	CommonParams

	BizId string `举例：	134523^4351232	发送回执ID,可根据该ID查询具体的发送状态`
}
