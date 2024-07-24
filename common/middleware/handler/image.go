package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path"
)

func ImageShow(c *gin.Context) {

	name := c.Param("name")

	file, readError := ioutil.ReadFile(path.Join("images", name))
	if readError != nil {
		c.JSON(200, gin.H{
			"message": "暂无图片",
		})
		return
	}
	c.Writer.WriteString(string(file))

}
