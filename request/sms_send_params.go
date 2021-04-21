package request

import (
	"encoding/json"
	"errors"
	"strings"
)

//短信发送需要的业务参数，查看文档：https://help.aliyun.com/document_detail/55284.html
type SmsSendParams struct {
	//Action   string `必填：是	API的命名，固定值，如发送短信API的值为：SendSms`
	//Version  string `必填：是	API的版本，固定值，如短信API的值为：2017-05-25`
	//RegionId string `必填：是	API支持的RegionID，如短信API的值为：cn-hangzhou`

	PhoneNumbers    string            `类型：String	必填：是 说明：15000000000	短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式`
	SignName        string            `类型：String	必填：是 说明：云通信	短信签名`
	TemplateCode    string            `类型：String	必填：是 说明：SMS_0000	短信模板ID`
	TemplateParam   map[string]string `类型：String	必填：否 说明：{“code”:”1234”,”product”:”ytx”}	短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败`
	SmsUpExtendCode string            `类型：String	必填：否 说明：90999	上行短信扩展码,无特殊需要此字段的用户请忽略此字段`
	OutId           string            `类型：String	必填：否 说明：abcdefgh	外部流水扩展字段`
}

func BuildSmsSendParams(phoneNumbers, templateCode, signName string, templateParam map[string]string) *SmsSendParams {
	return &SmsSendParams{
		PhoneNumbers:  phoneNumbers,
		TemplateParam: templateParam,
		TemplateCode:  templateCode,
		SignName:      signName,
	}
}

//实现BizParmas接口
func (ssp SmsSendParams) ToMap() (map[string]string, error) {
	if ssp.PhoneNumbers == "" {
		return nil, errors.New("SmsSendParams must set PhoneNumbers")
	}

	if len(strings.Split(ssp.PhoneNumbers, ",")) > 1000 {
		return nil, errors.New("SmsSendParams's Phone Numbers count can't more then 1000")
	}

	if ssp.TemplateCode == "" {
		return nil, errors.New("SmsSendParams must set TemplateCode")
	}

	if ssp.SignName == "" {
		return nil, errors.New("SmsSendParams must set SignName")
	}

	result := map[string]string{
		"Action":       "SendSms",
		"Version":      DEFAULT_VERSION,
		"RegionId":     DEFAULT_REGION_ID,
		"TemplateCode": ssp.TemplateCode,
		"PhoneNumbers": ssp.PhoneNumbers,
		"SignName":     ssp.SignName,
	}

	if ssp.SmsUpExtendCode != "" {
		result["smsUpExtendCode"] = ssp.SmsUpExtendCode
	}

	if ssp.OutId != "" {
		result["OutId"] = ssp.OutId
	}

	if ssp.TemplateParam != nil {
		jsonString, err := json.Marshal(ssp.TemplateParam)
		if err != nil {
			return nil, errors.New("SmsSendParams'TempalteParam to json fail")
		}
		result["TemplateParam"] = string(jsonString)
	}

	return result, nil
}
