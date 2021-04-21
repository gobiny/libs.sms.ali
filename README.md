### 调用方式
1. 获取本lib:
```
go get github.com/liuximu/libs.sms.ali
```

2. 发送短信 SendSms:
```
package main

const (
	ACCESS_KEY_ID     string = "your AccessKeyId"
	ACCESS_KEY_SECRET string = "your AccessKeySecret"
)

func main(){
	client := Create(ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	params := map[string]string{"code": "1111"}
	response, err := client.SendSms("手机号码，多个用,分隔", "短信模板（需要申请）", "短信签名（需要申请）", params)
}
```

3. 查看短信发送状态 QuerySendDetails:
```
package main

const (
	ACCESS_KEY_ID     string = "your AccessKeyId"
	ACCESS_KEY_SECRET string = "your AccessKeySecret"
)

func main(){
	client := Create(ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	params := map[string]string{"code": "1111"}
	resp, err := client.QuerySms("手机号码", "日期（比如20140225）", "BizId（可为空）", 50, 1)
}
```

### 起因
因为想使用阿里云(aliyun)的短信(SendSms)发送功能，但是官方没有提供Golang的SDK，GitHub上有找到几个，包括 [gwpp/alidayu-go](https://github.com/gwpp/alidayu-go)，但是阅读源码后感觉还是不清晰，所以自己写一个。

### 实现功能
这个项目实现了短信发送（SendSms）、短信发送状态查询（QuerySendDetails）和短信回执消息查询（SmsUp）三个功能，进行简单解释：
- SendSms：向单个|一系列用户发送短消息；[资料](https://help.aliyun.com/document_detail/55284.html)
- QuerySendDetails：查询向单个用户发送消息的状态；[资料](https://help.aliyun.com/document_detail/55289.html)
- SmsUp：查询用户回复的消息 [资料](https://help.aliyun.com/document_detail/55496.html) (未实现)


### 实现细节
[HTTP协议及签名](https://help.aliyun.com/document_detail/56189.html)是阿里云通信产品的API签名规范，短信服务就是其中一类。它是将数据按照[POP](https://help.aliyun.com/document_detail/54229.html)协议进行加密，然后发送http/https请求来完成服务的调用，那么整个流程就拆解为：

```
应用构造为未密请求参数

将请求参数加密

发送请求并返回响应

应用处理响应
```

中间两步是本类库的核心。

类库模块包括：
- 请求对象化模块(request/)
- 加密模块(utlis/)
- 响应对象化模块(response/)

