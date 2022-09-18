<div class="content">
    <header>{{.Post.Title}}</header>
    <main>{{.Post.Content}}</main>
    <footer>
        <div class="footer-leftbox">
            <p>
                Like: {{.Post.LikeCount}}
            </p>
            <p>
                ViewCount: {{.Post.view_count}}
            </p>
        </div>
    </footer>
</div>