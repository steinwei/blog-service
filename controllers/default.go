package controllers

import (
	"blog/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var post []models.Post
	o := orm.NewOrm().QueryTable(new(models.Post))
	o.Limit(3).All(&post)
	c.Data["post"] = post
	c.Data["title"] = "test"
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
