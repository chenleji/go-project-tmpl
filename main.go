package main

import (
	"os"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/chenleji/go-project-tmpl/routers"
	"github.com/chenleji/go-project-tmpl/helper"
)

func main() {
	registerService()
	beego.Run()
}

func registerService() {
	if beego.BConfig.RunMode == "prod" {
		err := helper.RegistryConsul("orchestration", 8080, "/health")
		if err != nil {
			logs.Error(err.Error())
			os.Exit(-1)
		}
	}
}
