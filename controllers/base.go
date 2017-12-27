package controllers

import (
	"errors"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

var (
	NoErr        = &ControllerError{200, "OK"}
	Err404       = &ControllerError{404, "page not found"}
	ErrInputData = &ControllerError{400, "数据输入错误"}
	ErrConnToMQ  = &ControllerError{500, "连接消息队列错误"}
	ErrSendToMQ  = &ControllerError{500, "发送消息到队列错误"}
	ErrDatabase  = &ControllerError{500, "数据库错误"}
	ErrCronExp   = &ControllerError{400, "Cron表达式格式不正确"}
)

type ControllerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseController struct {
	beego.Controller
}

func (base *BaseController) VerifyInputData(obj interface{}) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		return err
	}
	if !ok {
		str := ""
		for _, err := range valid.Errors {
			str += err.Key + ":" + err.Message + ";"
		}
		return errors.New(str)
	}

	return nil
}

func (base *BaseController) RespJson(obj interface{}) {
	base.Data["json"] = obj
	base.ServeJSON()
}

func (base *BaseController) RespError(ctlErr *ControllerError) {
	base.Ctx.Output.Header("content-type", "application/json")
	errBytes, _ := json.Marshal(ctlErr)
	base.CustomAbort(ctlErr.Code, string(errBytes))
}
