package main

import (
	"go-admin/cmd"
	"go-admin/common/global"
	_ "go-admin/global/initialization"
)

//go:generate swag init --parseDependency --parseDepth=6
var Version string
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	global.Version = Version
	cmd.Execute()

}
