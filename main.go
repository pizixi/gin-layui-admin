package main

import (
	"log"

	"gin-layui-admin/conf"
	"gin-layui-admin/models"
	"gin-layui-admin/router"

	"github.com/gin-gonic/gin"
)

// go build -ldflags "-s -w"
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	models.Init()

	r := gin.Default()
	router.RegRouter(r)
	appYaml := conf.YamlConf.App
	if appYaml.Debug {
		r.Run(":" + appYaml.Port)
	} else {
		r.Run("127.0.0.1:" + appYaml.Port)
	}
}
