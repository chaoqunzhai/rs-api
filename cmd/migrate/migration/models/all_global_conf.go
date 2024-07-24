/**
@Author: chaoqun
* @Date: 2023/12/19 17:44
*/
package models

//全局生效配置

type GlobalArticle struct {
	RichGlobal
	Name      string       `gorm:"size:50;comment:标题名称"`
	Subtitle string   `gorm:"size:120;comment:副标题名称"`
	Type int  `json:"type" gorm:"size:1;default:1;index;comment:1:公告 2:帮助文档"`
	Link string `json:"link" gorm:"size:120;comment:跳转地址"`

}
func (GlobalArticle) TableName() string {
		return "global_article"
}

type CompanyRegister struct {
	Model
	ModelTime
	UserName string `json:"user_name" gorm:"size:20;comment:名称"`
	Phone string `json:"phone" gorm:"size:11;comment:手机号"`
	Desc string `json:"desc" gorm:"size:100;comment:描述"`

}
func (CompanyRegister) TableName() string {
	return "company_register"
}


//全局消息通知 用于给大B的消息推送
type GlobalNotice struct {
	Model
	ControlBy
	ModelTime
	CId int `json:"c_id" gorm:"default:0;comment:为0时,大B都可以获取到,非0时,大B专属消息"`
	Layer  int    `json:"layer" gorm:"default:1;index;comment:排序"` //排序
	Enable bool   `json:"enable" gorm:"default:true;comment:开关"`
	Value string   `gorm:"size:120;comment:内容"`

}
func (GlobalNotice) TableName() string {
	return "global_notice"
}