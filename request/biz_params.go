package request

const (
	DEFAULT_REGION_ID = "cn-hangzhou"
	DEFAULT_VERSION   = "2017-05-25"
)

//业务参数接口，供其他子类实现
type BizParams interface {
	//将本结构体转换为map
	ToMap() (map[string]string, error)
}
