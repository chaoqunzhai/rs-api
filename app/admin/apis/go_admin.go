package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/common/global"
)

func GoAdmin(c *gin.Context) {
	c.String(200, fmt.Sprintf("rs-api %s", global.Version))
}
