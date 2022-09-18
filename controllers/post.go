package controllers

import (
    "blog/models"
    "blog/services"
    "strconv"

    "github.com/beego/beego/v2/client/orm"
    beego "github.com/beego/beego/v2/server/web"
)

type PostController struct {
    beego.Controller
}

func (c *PostController) Get() {
    id := c.Ctx.Input.Param(":id")
    postId, _ := strconv.ParseUint(id, 0, 8)
    post := models.Post{Id: uint(postId)}
    o := orm.NewOrm()
    o.QueryTable("post").Filter("id", uint(postId)).Update(orm.Params{
        "view_count": orm.ColValue(orm.ColAdd, 1),
    })

    if err := o.Read(&post); err != nil {
        c.Abort("404")
    }

    comments := services.GetComment(int(postId))
    c.Data["comments"] = comments
    c.Data["post"] = post
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["Comment"] = "comment.tpl"
    c.LayoutSections["Scripts"] = "main.js"
    c.Layout = "layout.tpl"
    c.TplName = "post.tpl"
}
