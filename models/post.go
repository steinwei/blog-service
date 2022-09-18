package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Post struct {
	Id           uint      `orm:"column(id);auto;size(11)" description:"table id" json:"id"`
	Title        string    `orm:"column(title);size(50)" description:"title" json:"title"`
	Introduction string    `orm:"column(introduction);size(200)" description:"description" json:"description"`
	Content      string    `orm:"column(content);type(text)" description:"content" json:"content"`
	LikeCount    uint      `orm:"column(like_count);type(text)" description:"likeCount" json:"like_count"`
	ViewCount    uint      `orm:"column(view_count);type(text)" description:"viewCount" json:"view_count"`
	CreatedAt    time.Time `orm:"type(datetime);column(created_at);" description:"createdAt" json:"created_at"`
	UpdatedAt    time.Time `orm:"type(datetime);column(updated_at);" description:"updatedAt" json:"updated_at"`
	PublishedAt  time.Time `orm:"type(datetime);column(published_at);" description:"publishedAt" json:"published_at"`
}

func init() {
	orm.RegisterModel(new(Post))
}
