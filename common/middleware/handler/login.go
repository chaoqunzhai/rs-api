package handler

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"go-admin/global"
	"gorm.io/gorm"
)

type Login struct {
	Phone    string `form:"phone" json:"phone" `
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"`
	CID      string `form:"cid" json:"cid" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
}

func (u *Login) GetUserPhone(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("phone = ?  and status = ? and enable = ?",
		u.Phone, global.SysUserSuccess, true).Limit(1).First(&user).Error
	if err != nil {
		err = errors.New("手机号不存在")
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		err = errors.New("手机号或者密码错误")
		return
	}

	return
}
func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = '2'", u.UserName).First(&user).Error
	if err != nil {
		err = errors.New("用户不存在")
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		err = errors.New("用户名或密码错误")
		return
	}

	return
}
