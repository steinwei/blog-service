package models

import (
    "time"

    "github.com/beego/beego/v2/client/orm"
)

type Comment struct {
    Id        uint      `orm:"column(id);auto;size(11)" description:"table_id"`
    PostId    int       `orm:"column(postid);size(10);default(0)" description:"post_id"`
    Username  string    `orm:"column(username);size(30);default('')" description:"user_id"`
    ParentId  int       `orm:"column(parentid);size(10);default(0)" description:"parent_id"`
    Content   string    `orm:"column(content);type(text)" description:"content"`
    CreatedAt time.Time `orm:"auto_now;column(created_at);type(datetime)" description:"created_at"`
}

func init() {
    orm.RegisterModel(new(Comment))
}
