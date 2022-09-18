package services

import (
    "blog/models"

    "github.com/beego/beego/v2/client/orm"
)

type CommentService struct{}

type Comments struct {
    Id        int         `json:"id"`
    PostId    int         `json:"post_id"`
    ParentId  int         `json:"parent_id"`
    Username  string      `json:"username"`
    Content   string      `json:"content"`
    CreatedAt string      `json:"created_at"`
    Children  []*Comments `json:"children"`
}

func GetComment(postId int) []*Comments {
    var comments []models.Comment
    o := orm.NewOrm().QueryTable(new(models.Comment))
    o.Filter("postId", postId).OrderBy("-created_at").All(&comments)
    tree := BuildTree(comments, 0)
    return tree
}

func BuildTree(comments []models.Comment, parentId int) []*Comments {
    c := make([]*Comments, 0)
    for _, v := range comments {
        var r = new(Comments)
        r.Id = int(v.Id)
        r.ParentId = v.ParentId
        r.PostId = v.PostId
        r.Username = v.Username
        r.Content = v.Content
        r.CreatedAt = v.CreatedAt.Format("2022-01-01")
        pId := v.ParentId
        if pId == parentId {
            Id := int(v.Id)
            children := BuildTree(comments, Id)
            if len(children) > 0 {
                r.Children = BuildTree(comments, Id)
            }
            c = append(c, r)
        }
    }
    return c
}
