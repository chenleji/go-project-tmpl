package routers

import (
	"github.com/chenleji/go-project-tmpl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthCheckController{})
}
