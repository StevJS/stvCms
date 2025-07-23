package request

type CreatePostRequest struct {
	Title   string `json:"title" required:"true"`
	Content string `json:"content" required:"true"`
	Author  string `json:"author" required:"true"`
	//Images  []byte `json:"images"`
}

type CodeContent struct {
	PostID  int     `json:"post_id" required:"true"`
	Content Content `json:"content" required:"true"`
}

type Content struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

type UpdatePostRequest struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
