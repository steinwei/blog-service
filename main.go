package main

import (
	_ "blog/routers"

	_ "blog/models"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

