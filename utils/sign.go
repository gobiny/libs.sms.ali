package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

//使用HmacSHA1加密字符串
func HmacSha1Sign(planText, key string) string {
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(planText))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return specialUrlEncode(sign)
}
