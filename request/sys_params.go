package request

import (
	"errors"
	"time"

	"github.com/liuximu/libs.sms.ali/utils"
)

const (
	SP_FORMAT_JSON string = "JSON"
	SP_FORMAT_XML  string = "XML"
)

//POP加密中的基本参数，这些参数可以认为是业务层不需要关心的，系统会自动你个填充
type SysParams struct {
	AccessKeyId      string `必填：是`
	Timestamp        string `必填：是	格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT`
	Format           string `必填：否	没传默认为JSON，可选填值：XML`
	SignatureMethod  string `必填：是	建议固定值：HMAC-SHA1`
	SignatureVersion string `必填：是	建议固定值：1.0`
	SignatureNonce   string `必填：是	用于请求的防重放攻击，每次请求唯一`
	Signature        string `必填：是	最终生成的签名结果值`
}

//实现BizParmas接口
func (sp SysParams) ToMap() (map[string]string, error) {
	if sp.AccessKeyId == "" {
		return nil, errors.New("SysParams must set AccessKeyId")
	}
	result := make(map[string]string)
	result["AccessKeyId"] = sp.AccessKeyId

	if sp.Timestamp == "" {
		sp.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z") //2017-08-30T18:22:15Z
	}
	result["Timestamp"] = sp.Timestamp

	if sp.Format == "" {
		sp.Format = SP_FORMAT_JSON
	}
	result["Format"] = sp.Format

	if sp.SignatureMethod == "" {
		sp.SignatureMethod = "HMAC-SHA1"
	}
	result["SignatureMethod"] = sp.SignatureMethod

	if sp.SignatureVersion == "" {
		sp.SignatureVersion = "1.0"
	}
	result["SignatureVersion"] = sp.SignatureVersion

	if sp.SignatureNonce == "" {
		sp.SignatureNonce = utils.GetUid()
	}
	result["SignatureNonce"] = sp.SignatureNonce

	return result, nil
}
