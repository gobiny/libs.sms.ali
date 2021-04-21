package response

/**
{
	"TotalCount": 3,
	"Message": "OK",
	"RequestId": "CE0B53F3-EF4B-45B4-8D02-402E1C22D717",
	"SmsSendDetailDTOs": {
		"SmsSendDetailDTO": [
			{
				"SendDate": "2017-08-31 18:00:17",
				"SendStatus": 3,
				"ReceiveDate": "2017-08-31 18:00:21",
				"ErrCode": "0",
				"TemplateCode": "SMS_90120021",
				"Content": "【熙穆】你的验证码是1111，有效时间3分钟。若本人操作请及时输入，否则请忽略。",
				"PhoneNum": "18579051698"
			},
			{
				"SendDate": "2017-08-31 12:37:47",
				"SendStatus": 3,
				"ReceiveDate": "2017-08-31 12:37:56",
				"ErrCode": "0",
				"TemplateCode": "SMS_90120021",
				"Content": "【熙穆】你的验证码是1111，有效时间3分钟。若本人操作请及时输入，否则请忽略。",
				"PhoneNum": "18579051698"
			}
		]
	},
	"Code": "OK"
}
*/
//参考文档：https://help.aliyun.com/document_detail/55452.html
type SmsQueryParams struct {
	CommonParams

	TotalCount        int //100	发送总条数
	TotalPage         int //10	总页数
	SmsSendDetailDTOs SmsSendDetailDTOs
}

type SmsSendDetailDTOs struct {
	SmsSendDetailDTO []SmsSendDetailDTO
}

type SmsSendDetailDTO struct {
	SendDate     string
	SendStatus   int //发送状态 1：等待回执，2：发送失败，3：发送成功
	ReceiveDate  string
	ErrCode      string
	TemplateCode string
	Content      string
	PhoneNum     string
}
