package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/chenleji/go-project-tmpl/helper"
)

func init() {
	orm.Debug = false
	e := helper.EnvVar{}.Get()

	dbURL := e.DbUrl
	dbPort := e.DbPort
	dbUser := e.DbUser
	dbPassword := e.DbPwd
	dbName := "model"

	if beego.BConfig.RunMode == "dev" {
		dbURL = "localhost"
		dbPort = "3306"
		dbUser = "root"
		dbPassword = ""
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbURL, dbPort, dbName)
	maxIdle := 30
	maxConn := 50
	logs.Debug("dataSource:", dataSource)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	//orm.RegisterModel(new(Model))
	orm.RunSyncdb("default", false, true)
}
