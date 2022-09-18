<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="../static/css/layout.css">
    <title>{{.Title}}</title>
    {{.HtmlHead}}
</head>
<body>
    <section id="main-container">
        <header class="nav">
            <div class="logo">slot</div>
            <ul>
                {{range $key, $item := .navs}}
                <li class="nav-item">
                    <div>
                        <a href="#">{{$item.name}}</a>
                    </div>
                </li>
                {{end}}
            </ul>
        </header>
        <main>
            {{.LayoutContent}}
            {{.Comment}}
        </main>
        <footer></footer>
    </section>
    <script src="../static/js/reload.min.js"></script>
    <script src="../static/js/layout.js"></script>
</body>
</html>