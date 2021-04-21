package sms

import (
	"encoding/json"

	"github.com/liuximu/libs.sms.ali/request"
	"github.com/liuximu/libs.sms.ali/response"
	"github.com/liuximu/libs.sms.ali/utils"
)

//客户端，实现根据配置进行请求的发送
type Client struct {
	SysParams       *request.SysParams
	accessKeySecret string
}

//创建一个客户端
func Create(accessKeyId, accessKeySecret string) *Client {
	sp := &request.SysParams{AccessKeyId: accessKeyId}
	return CreateWithSysParams(sp, accessKeySecret)
}

func CreateWithSysParams(sp *request.SysParams, accessKeySecret string) *Client {
	return &Client{
		SysParams:       sp,
		accessKeySecret: accessKeySecret,
	}
}

type IClient interface {
	SendSms(phoneNumbers, templateCode, signName string, templateParmas map[string]string) (response.SmsSendParams, error)
	QuerySms(phoneNumber, sendDate, bizId string, pageSize, curPage int) (response.SmsQueryParams, error)
	Execute(request.BizParams) ([]byte, error)
}

//phoneNumbers 机号码，多个用,分隔
//templateCode 短信模板 （需要申请）
//signName 短信签名（需要申请）
//templateParams 短信内容参数
func (client *Client) SendSms(phoneNumbers, templateCode, signName string, templateParmas map[string]string) (resObj response.SmsSendParams, err error) {
	bp := request.BuildSmsSendParams(phoneNumbers, templateCode, signName, templateParmas)
	resp, err := client.Execute(bp)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &resObj)
	return
}

//phoneNumber 要查询的手机号
//sendDate 日期，比如 20170502
//bizId 可以为空
//pageSize 每页数量 1 - 50
//curPage 当前页 1 开始
func (client *Client) QuerySms(phoneNumber, sendDate, bizId string, pageSize, curPage int) (resObj response.SmsQueryParams, err error) {
	sqp := request.BuildSmsQueryParams(phoneNumber, sendDate, bizId, pageSize, curPage)
	resp, err := client.Execute(sqp)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &resObj)
	return
}

func (client *Client) Execute(bp request.BizParams) (response []byte, err error) {
	//合并参数
	params, err := bp.ToMap()
	if err != nil {
		return
	}
	spMap, err := client.SysParams.ToMap()
	if err != nil {
		return
	}

	for k, v := range spMap {
		params[k] = v
	}

	//构造url
	url := utils.BuildUrl(client.accessKeySecret, &params)
	//panic(url)

	//发送请求
	response, err = utils.HttpGet(url)
	return
}
