package models

type RsMenu struct {
	Model
	ModelTime
	Layer     int    `gorm:"size:5;index;comment:排序"` //排序
	Enable    bool   `gorm:"comment:开关"`
	Role      string `gorm:"size:20;comment:哪个角色菜单"`
	Name      string `gorm:"size:30;comment:英文名称"`
	Path      string `gorm:"size:30;comment:路径,也是权限名称"`
	ParentId  int    `json:"parentId" gorm:"index;size:11;comment:父ID"`
	MetaTitle string `gorm:"size:30;comment:标题"`
	MetaIcon  string `gorm:"size:30;comment:图片"`
	Hidden    bool   `gorm:"comment:是否隐藏"`
	KeepAlive bool   `gorm:"comment:是否缓存"`
	Component string `gorm:"size:50;comment:import路径"`
}

func (RsMenu) TableName() string {
	return "rs_menu"
}
