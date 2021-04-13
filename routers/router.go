package routers

import (
	"github.com/beego/beego/v2/server/web"

	"boilerplate-golang/controllers"
)

func init() {

	web.Router("/", &controllers.MainController{}, "get:Get")

	web.Router("/users", &controllers.UserController{}, "get:GetUsers")
	web.Router("/user/id", &controllers.UserController{}, "get:GetUserById")
	web.Router("/user/username", &controllers.UserController{}, "get:GetUserByUsername")
	web.Router("/user/id", &controllers.UserController{}, "put:UpdateUserById")
}
