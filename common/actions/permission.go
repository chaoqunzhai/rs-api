package actions

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	sys "go-admin/app/admin/models"
	"go-admin/global"
	"gorm.io/gorm"
)

type DataPermission struct {
	DataScope int
	UserId    int
	DeptId    int
	RoleId    int
	Enable    bool
	CId       int
}

type UserDataPermission struct {
	UserId    int
	CId       int
}
//超管平台 必须都是超管权限才可以访问
func PermissionAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		var p = new(DataPermission)
		if userId := user.GetUserIdStr(c); userId != "" {
			p, err = newDataPermission(db, userId)
			if err != nil {
				log.Errorf("MsgID[%s] PermissionAction error: %s", msgID, err)
				response.Error(c, 500, err, "权限范围鉴定错误")
				c.Abort()
				return
			}
		}
		if !p.Enable {
			response.Error(c, 401, errors.New("您账户已被停用！"), "您账户已被停用！")
			c.Abort()
			return
		}

		if p.RoleId == 0 {
			response.Error(c, 401, errors.New("您没有权限访问"), "您没有权限访问")
			c.Abort()
			return
		}
		//权限校验
		if p.DataScope != global.RoleSuper {
			response.Error(c, 401, errors.New("您没有权限访问"), "您没有权限访问")
			c.Abort()
			return
		}

		c.Set(PermissionKey, p)
		c.Next()
	}
}
func init() {

}


func newDataPermission(tx *gorm.DB, userId interface{}) (*DataPermission, error) {
	var err error
	p := &DataPermission{}

	err = tx.Table("sys_user").
		Select("sys_user.user_id", "sys_role.role_id", "sys_user.c_id", "sys_user.enable", "sys_role.data_scope").
		Joins("left join sys_role on sys_role.data_scope = sys_user.role_id").
		Where("sys_user.user_id = ?", userId).
		Scan(p).Error
	if err != nil {
		err = errors.New("获取用户数据出错 msg:" + err.Error())
		return nil, err
	}
	return p, nil
}

//使用模型生成的代码查询,使用gin的路由中间件查询到的数据(Context的Key）进行数据分层
func Permission(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		fmt.Printf("Permission:DB层的数据隔离,用户:%v,角色:%v\n",p.UserId,p.RoleId)
		if p.CId > 0 {
			return db.Table(tableName).Where("c_id = ?", p.CId)
		}else {
			//如果没有获取到用户的站点,那就获取默认的站点,防止数据泄露
			return db.Table(tableName).Where("c_id = 1", p.CId)
		}
	}
}


func PermissionSysUser(tableName string,p *sys.SysUser) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fmt.Printf("PermissionSysUser:DB层的数据隔离,用户:%v,站点:%v\n",p.UserId,p.CId)
		if p.CId > 0 {
			return db.Table(tableName).Where("c_id = ?", p.CId)
		}else {
			//如果没有获取到用户的站点,那就获取默认的站点,防止数据泄露
			return db.Table(tableName).Where("c_id = 1", p.CId)
		}
	}
}

func getPermissionFromContext(c *gin.Context) *DataPermission {
	p := new(DataPermission)
	if pm, ok := c.Get(PermissionKey); ok {
		switch pm.(type) {
		case *DataPermission:
			p = pm.(*DataPermission)
		}
	}
	return p
}

// GetPermissionFromContext 提供非action写法数据范围约束
func GetPermissionFromContext(c *gin.Context) *DataPermission {
	return getPermissionFromContext(c)
}
