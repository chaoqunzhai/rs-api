/*
*
@Author: chaoqun
* @Date: 2024/5/8 00:24
*/
package api

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
	"strconv"
)

// 系统操作权限
func SystemPermission() []map[string]string {
	list := make([]map[string]string, 0)
	list = append(list, map[string]string{
		"name":       "/system/shop_config",
		"layer":      "928",
		"meta_title": "店铺设置",
	})
	list = append(list, map[string]string{
		"name":       "/system/user_manager",
		"layer":      "927",
		"meta_title": "员工管理",
	})
	list = append(list, map[string]string{
		"name":       "/system/user/list",
		"layer":      "926",
		"parent":     "/system/user_manager",
		"meta_title": "员工列表",
	})
	list = append(list, map[string]string{
		"name":       "/system/user/add",
		"layer":      "925",
		"parent":     "/system/user_manager",
		"meta_title": "员工创建",
	})
	list = append(list, map[string]string{
		"name":       "/system/user/update",
		"layer":      "924",
		"parent":     "/system/user_manager",
		"meta_title": "员工更新",
	})
	list = append(list, map[string]string{
		"name":       "/system/user/remove",
		"layer":      "923",
		"parent":     "/system/user_manager",
		"meta_title": "员工删除",
	})
	list = append(list, map[string]string{
		"name":       "/system/user/update_password",
		"layer":      "922",
		"parent":     "/system/user_manager",
		"meta_title": "员工密码重置",
	})
	list = append(list, map[string]string{
		"name":       "/system/update_password",
		"layer":      "921",
		"meta_title": "账户密码修改",
	})
	return list
}
func createRow(db *gorm.DB, row map[string]string) {
	layer, _ := strconv.Atoi(row["layer"])
	object := models.RsMenu{
		Name:      row["name"],
		Enable:    true,
		MetaTitle: row["meta_title"],
		Layer:     layer,
	}
	var parentObject models.RsMenu
	parentName := row["parent"]
	if parentName != "" {
		db.Model(&models.RsMenu{}).Select("id").Where("name = ? and enable = ?", parentName, true).Limit(1).Find(&parentObject)

	}

	//如果存在 除了name是不更新 其他都需要更新
	var DyMbmMenu models.RsMenu
	db.Model(&models.RsMenu{}).Where("name = ?", row["name"]).Limit(1).Find(&DyMbmMenu)
	if DyMbmMenu.Id > 0 {
		DyMbmMenu.MetaTitle = row["meta_title"]
		DyMbmMenu.Layer = layer
		DyMbmMenu.ParentId = parentObject.Id
		DyMbmMenu.Enable = true
		db.Save(&DyMbmMenu)
	} else {
		object.ParentId = parentObject.Id
		db.Create(&object)
	}
}
func InitializationMbmWeApp() {
	dbs := sdk.Runtime.GetDb()
	for _, db := range dbs {

		for _, row := range SystemPermission() {

			createRow(db, row)

		}
	}
}
