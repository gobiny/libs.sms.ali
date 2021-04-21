package request

import (
	"errors"
	"strconv"
)

type SmsQueryParams struct {
	PhoneNumber string //必须	15000000000	短信接收号码,如果需要查询国际短信,号码前需要带上对应国家的区号,区号的获取详见国际短信支持国家信息查询API接口
	SendDate    string //必须	20170525	短信发送日期格式yyyyMMdd,支持最近30天记录查询
	PageSize    int    //必须	10	页大小Max=50
	CurrentPage int    //必须	1	当前页码
	BizId       string //可选	1234^1234	发送流水号,从调用发送接口返回值中获取
}

func BuildSmsQueryParams(phoneNumber, sendDate, bizId string, pageSize, curPage int) *SmsQueryParams {
	return &SmsQueryParams{
		PhoneNumber: phoneNumber,
		SendDate:    sendDate,
		PageSize:    pageSize,
		CurrentPage: curPage,
		BizId:       bizId,
	}
}

func (sqp SmsQueryParams) ToMap() (map[string]string, error) {
	if sqp.PhoneNumber == "" {
		return nil, errors.New("SmsQueryParams must set PhoneNumber")
	}

	if sqp.SendDate == "" {
		return nil, errors.New("SmsQueryParams must set SendDate")
	}

	if sqp.PageSize < 1 || sqp.PageSize > 50 {
		return nil, errors.New("SmsQueryParams's PageSize in [1,50]")
	}

	if sqp.CurrentPage < 1 {
		return nil, errors.New("SmsQueryParams's CurrentPage > 0")
	}

	params := map[string]string{
		"Action":      "QuerySendDetails",
		"Version":     DEFAULT_VERSION,
		"RegionId":    DEFAULT_REGION_ID,
		"PhoneNumber": sqp.PhoneNumber,
		"SendDate":    sqp.SendDate,
		"PageSize":    strconv.Itoa(sqp.PageSize),
		"CurrentPage": strconv.Itoa(sqp.CurrentPage),
	}

	if sqp.BizId != "" {
		params["BizId"] = sqp.BizId
	}

	return params, nil
}
