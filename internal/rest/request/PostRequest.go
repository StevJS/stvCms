package request

type CreatePostRequest struct {
	Title   string `json:"title" required:"true"`
	Content string `json:"content" required:"true"`
	Author  string `json:"author" required:"true"`
	Images  []byte `json:"images"`
}
