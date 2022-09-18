<!DOCTYPE html>

<html>
<head>
  <title>blog</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    here is header.
  </header>
  {{range $key,$post := .post}}
      <ul>
        <li>{{$post.Title}}</li>
        <li>{{$post.Content}}</li>
        <li>views: {{$post.ViewCount}}</li>
        <li>likes: {{$post.LikeCount}}</li>
      </ul>
  {{end}}
  <footer>
    here is footer.
  </footer>
</body>
</html>
