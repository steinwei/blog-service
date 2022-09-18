package main

import (
	_ "blog/inits"
	_ "blog/routers"

	beego "github.com/beego/beego"
	"github.com/beego/beego/orm"
)

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
