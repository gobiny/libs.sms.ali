package utils

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const URL string = "http://dysmsapi.aliyuncs.com/"

//查看文档：
//	https://help.aliyun.com/document_detail/54229.html
//	https://help.aliyun.com/document_detail/56189.html
// 返回完整的URL
func BuildUrl(key string, params *map[string]string) string {
	//将参数按照key去排序
	sortedParams := sortParams(params)
	//将排序了的参数拼成字符串
	url := concatQueryParams(params, sortedParams)
	//得到待签名的明文
	plainText := getSigningText(url)
	//得到明文的签名
	key = fmt.Sprintf("%s%s", key, "&")
	sign := HmacSha1Sign(plainText, key)

	return fmt.Sprintf("%s?Signature=%s&%s", URL, sign, url)
}

func sortParams(params *map[string]string) []string {
	keys := []string{}
	for k, _ := range *params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func concatQueryParams(params *map[string]string, sortedKey []string) string {
	input := ""
	for _, k := range sortedKey {
		input = fmt.Sprintf("%s&%s=%s", input, specialUrlEncode(k), specialUrlEncode((*params)[k]))
	}
	return input[1:]
}

func getSigningText(url string) string {
	return fmt.Sprintf("%s%s", "GET&%2F&", specialUrlEncode(url))
}

func specialUrlEncode(input string) string {
	input = url.QueryEscape(input)
	input = strings.Replace(input, "+", "%20", -1)
	input = strings.Replace(input, "*", "%2A", -1)
	return strings.Replace(input, "%7E", "~", -1)
}
