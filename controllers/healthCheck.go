package controllers

type HealthCheckController struct {
	BaseController
}

func (c *HealthCheckController) Get() {
	c.Data["json"] = map[string]interface{}{
		"active": true,
	}

	c.ServeJSON()
}
