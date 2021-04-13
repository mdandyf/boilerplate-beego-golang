package main

import (
	_ "boilerplate-golang/routers"
	"boilerplate-golang/utilities"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	utilities.GetDB()

	web.Run()
}
