package services

import (
	"stvCms/internal/models"
	"stvCms/internal/repository"
	"stvCms/internal/rest/request"
	"stvCms/internal/rest/response"
)

type IPostService interface {
	CreatePost(req request.CreatePostRequest) (string, error)
	GetPosts() ([]response.PostResponse, error)
}

type postService struct {
	repository repository.IPostRepository
}

func NewPostService() IPostService {
	return &postService{
		repository: repository.NewPostGormRepository(),
	}
}

func (ps *postService) CreatePost(req request.CreatePostRequest) (string, error) {
	post := reqToModel(req)

	response, err := ps.repository.CreatePost(post)
	if err != nil {
		return "No se pudo crear el post", err
	}

	return response, nil
}

func (ps *postService) GetPosts() ([]response.PostResponse, error) {

	posts := []response.PostResponse{}

	modelPosts, err := ps.repository.GetPosts()

	if err != nil {
		return posts, nil
	}

	for _, post := range modelPosts {
		data := response.PostResponse{
			Title:   post.Title,
			Content: post.Content,
			Author:  post.Author,
		}
		posts = append(posts, data)
	}

	return posts, nil
}

func reqToModel(req request.CreatePostRequest) models.Post {
	return models.Post{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		//Images:  req.Images,
	}
}
