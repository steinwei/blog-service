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

func (ctx *PostController) Get() {
    id := ctx.Ctx.Input.Param(":id")
    postId, _ := strconv.ParseUint(id, 0, 8)
    post := models.Post{Id: uint(postId)}
    o := orm.NewOrm()
    o.QueryTable("post").Filter("id", uint(postId)).Update(orm.Params{
        "view_count": orm.ColValue(orm.ColAdd, 1),
    })

    if err := o.Read(&post); err != nil {
        ctx.Abort("404")
    }

    comments := services.GetComment(int(postId))
    ctx.Data["comments"] = comments
    ctx.Data["post"] = post
    ctx.LayoutSections = make(map[string]string)
    ctx.LayoutSections["Comment"] = "comment.tpl"
    ctx.LayoutSections["Scripts"] = "main.js"
    ctx.Layout = "layout.tpl"
    ctx.TplName = "post.tpl"
}
