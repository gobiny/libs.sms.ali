package sms

import (
	"testing"
)

const (
	ACCESS_KEY_ID     string = "your AccessKeyId"
	ACCESS_KEY_SECRET string = "your AccessKeySecret"
)

func TestSendSms(t *testing.T) {
	client := Create(ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	params := map[string]string{"code": "1111"}
	response, err := client.SendSms("手机号码，多个用,分隔", "短信模板（需要申请）", "短信签名（需要申请）", params)
	if err != nil {
		t.Errorf("%v", err)
	}

	if response.Code != "OK" {
		t.Errorf("%v", response)
	}
}

func TestQuerySms(t *testing.T) {
	client := Create(ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	resp, err := client.QuerySms("手机号码", "日期（比如20140225）", "BizId（可为空）", 50, 1)
	if err != nil {
		t.Error(resp, err)
	}

}
