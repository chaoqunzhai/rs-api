package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	Work        string `json:"work"`
	AMap        AMap
	Redis       Redis  `json:"redis"`
	AdminKey string `json:"adminKey"`
	ImageUrl string  `json:"imageUrl"`
	Domain      string        `json:"domain"`
	ImageBase string        `json:"imageBase"`
	H5Url string `json:"h5Url"`
	Qiniu Qiniu `json:"qiniu"`
	Tx Tx `json:"tx"`
}
type Qiniu struct {
	AccessKey string `json:"AccessKey"`
	SecretKey string `json:"SecretKey"`
}
type Tx struct {
	CosSecretID string `json:"cosSecretID"`
	CosSecretKey string `json:"cosSecretKey"`
}
type Redis struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Password string `json:"password"`
}


type AMap struct {
	Key string
}
