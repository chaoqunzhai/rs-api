/*
*
@Author: chaoqun
* @Date: 2022/7/27 10:12
*/
package api

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
)

type MenusRow struct {
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	MetaTitle string     `json:"meta_title"`
	KeepAlive bool       `json:"keep_alive"`
	MetaIcon  string     `json:"meta_icon"`
	Hidden    bool       `json:"hidden"`
	Enable    bool       `json:"enable"`
	Layer     int        `json:"layer"`
	Role      string     `json:"role"`
	Children  []Children `json:"children"`
}
type Children struct {
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	MetaTitle string     `json:"meta_title"`
	KeepAlive bool       `json:"keep_alive"`
	MetaIcon  string     `json:"meta_icon"`
	Hidden    bool       `json:"hidden"`
	Enable    bool       `json:"enable"`
	Layer     int        `json:"layer"`
	Role      string     `json:"role"`
	Children  []Children `json:"children"`
}

const (
	RoleStr = "admin,company"
)

// layer 从小到大展示
// name就是前段的权限名称
func InitializationMenus() []MenusRow {
	var Menus []MenusRow
	ManageRenew := MenusRow{
		Name:      "/manage/renew",
		Path:      "/manage/renew",
		MetaTitle: "账户设置",
		Layer:     -1,
		MetaIcon:  "Icons.home",
	}
	Menus = append(Menus, ManageRenew)
	index := MenusRow{
		Name:      "/index",
		Path:      "/index",
		Component: "@/views/index/Index",
		MetaTitle: "概况",
		Layer:     0,
		MetaIcon:  "Icons.home",
	}
	Menus = append(Menus, index)
	report := MenusRow{
		Name:      "/report",
		Path:      "/report",
		Component: "@/views/report/Index",
		MetaTitle: "配送",
		Layer:     1,
		MetaIcon:  "Icons.give",
	}
	Menus = append(Menus, report)

	workerManager := MenusRow{
		Name:      "/worker/Download",
		Path:      "/worker/Download",
		MetaTitle: "下载",
		MetaIcon:  "Icons.market",
		Layer:     2,
	}

	orderExportManager := Children{
		Name:      "/worker/Order",
		Path:      "/worker/Order",
		MetaTitle: "订单导出",
		Layer:     3,
	}
	orderExportManager.Children = append(orderExportManager.Children, Children{
		Name:      "/worker/Report",
		Path:      "/worker/Report",
		Layer:     4,
		Component: "@/views/worker/Order/Report",
		MetaTitle: "配送报表",
		Hidden:    false,
	})
	orderExportManager.Children = append(orderExportManager.Children, Children{
		Name:      "/worker/Export",
		Layer:     5,
		Path:      "/worker/Export",
		Component: "@/views/worker/Order/Export",
		MetaTitle: "配送订单",
		Hidden:    false,
	})
	workerManager.Children = append(workerManager.Children, orderExportManager)
	Menus = append(Menus, workerManager)




	return Menus
}

func InsertChildrenDataMenus(parentId int, db *gorm.DB, menusData []Children) {

	for _, childrenRow := range menusData {

		var DyNaomiMenu models.RsMenu
		var childrenRowThisId int
		db.Model(&models.RsMenu{}).Select("id").Where("name = ?", childrenRow.Name).Limit(1).Find(&DyNaomiMenu)
		if DyNaomiMenu.Id > 0 {
			childrenRowThisId = DyNaomiMenu.Id
			db.Model(&models.RsMenu{}).Where("name = ?", childrenRow.Name).Updates(map[string]interface{}{
				"path":       childrenRow.Path,
				"meta_title": childrenRow.MetaTitle,
				"meta_icon":  childrenRow.MetaIcon,
				"keep_alive": childrenRow.KeepAlive,
				"hidden":     childrenRow.Hidden,
				"layer":      childrenRow.Layer,
			})
		} else {
			dto := &models.RsMenu{
				Name:      childrenRow.Name,
				Path:      childrenRow.Path,
				MetaTitle: childrenRow.MetaTitle,
				MetaIcon:  childrenRow.MetaIcon,
				Enable:    true,
				KeepAlive: childrenRow.KeepAlive,
				Hidden:    childrenRow.Hidden,
				Layer:     childrenRow.Layer,
				ParentId:  parentId,
			}
			db.Create(&dto)

			childrenRowThisId = dto.Id
		}

		if len(childrenRow.Children) > 0 {
			InsertChildrenDataMenus(childrenRowThisId, db, childrenRow.Children)
		}
	}

}

func Initialization() {

	fmt.Println("开始录入系统初始化配置")
	dbs := sdk.Runtime.GetDb()
	Menus := InitializationMenus()
	for _, db := range dbs {

		for _, menusData := range Menus {

			var DyNaomiMenu models.RsMenu
			var DyNaomiMenuId int
			db.Model(&models.RsMenu{}).Select("id").Where("name = ?", menusData.Name).Limit(1).Find(&DyNaomiMenu)
			if DyNaomiMenu.Id > 0 {
				DyNaomiMenuId = DyNaomiMenu.Id
				db.Model(&models.RsMenu{}).Where("name = ?", menusData.Name).Updates(map[string]interface{}{
					"path":       menusData.Path,
					"meta_title": menusData.MetaTitle,
					"meta_icon":  menusData.MetaIcon,
					"keep_alive": menusData.KeepAlive,
					"hidden":     menusData.Hidden,
					"layer":      menusData.Layer,
				})
			} else {
				row := &models.RsMenu{
					Name:      menusData.Name,
					Path:      menusData.Path,
					MetaTitle: menusData.MetaTitle,
					MetaIcon:  menusData.MetaIcon,
					Enable:    true,
					KeepAlive: menusData.KeepAlive,
					Hidden:    menusData.Hidden,
					Layer:     menusData.Layer,
				}
				db.Create(&row)
				DyNaomiMenuId = row.Id
			}

			if len(menusData.Children) > 0 {
				InsertChildrenDataMenus(DyNaomiMenuId, db, menusData.Children)
			}
		}

	}
}
