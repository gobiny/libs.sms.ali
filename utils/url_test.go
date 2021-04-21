package utils

import "testing"

func TestPopSign(t *testing.T) {
	m := &map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   "45e25e9b-0a6f-4070-8c85-2956eda1b466",
		"AccessKeyId":      "testId",
		"SignatureVersion": "1.0",
		"Timestamp":        "2017-07-12T02:42:19Z",
		"Format":           "XML",
		"Action":           "SendSms",
		"Version":          "2017-05-25",
		"RegionId":         "cn-hangzhou",
		"PhoneNumbers":     "15300000001",
		"SignName":         "阿里云短信测试专用",
		"TemplateParam":    `{"customer":"test"}`,
		"TemplateCode":     "SMS_71390007",
		"OutId":            "123",
	}
	url := BuildUrl("testSecret", m)
	except := "http://dysmsapi.aliyuncs.com/?Signature=zJDF%2BLrzhj%2FThnlvIToysFRq6t4%3D&AccessKeyId=testId&Action=SendSms&Format=XML&OutId=123&PhoneNumbers=15300000001&RegionId=cn-hangzhou&SignName=%E9%98%BF%E9%87%8C%E4%BA%91%E7%9F%AD%E4%BF%A1%E6%B5%8B%E8%AF%95%E4%B8%93%E7%94%A8&SignatureMethod=HMAC-SHA1&SignatureNonce=45e25e9b-0a6f-4070-8c85-2956eda1b466&SignatureVersion=1.0&TemplateCode=SMS_71390007&TemplateParam=%7B%22customer%22%3A%22test%22%7D&Timestamp=2017-07-12T02%3A42%3A19Z&Version=2017-05-25"
	if url != except {
		t.Error(url)
	}
}
