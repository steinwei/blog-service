package validations

type CommentValid struct {
    Name     string `valid:"Required;"`
    Email    string `valid:"Required;Email;"`
    Content  string `valid:"Required;"`
    PostId   int    `valid:"Required;"`
    ParentId int    `valid:"Required;"`
}
