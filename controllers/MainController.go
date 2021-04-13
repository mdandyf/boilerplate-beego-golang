package controllers

import (
	"boilerplate-golang/services"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

// Get ...
func (t *MainController) Get() {
	data, _ := services.GetMainService()

	t.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	t.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	t.Data["json"] = data
	t.ServeJSON()
}
