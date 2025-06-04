package services

import (
	"stvCms/internal/config"
	"stvCms/internal/models"
	"stvCms/internal/repository"
	"stvCms/internal/rest/request"
)

type IPostService interface {
	CreatePost(req request.CreatePostRequest) (string, error)
}

type postService struct {
	repository repository.IPostRepository
}

func NewPostService() IPostService {
	return &postService{
		repository: config.NewPostgresGormRepo
	}
}

func (ps *postService) CreatePost(req request.CreatePostRequest) (string, error) {
	post := reqToModel(req)

	response, err := ps.repository.CreatePost(post)
	if err != nil {
		return "No se pudo crear el post", err.Error()
	}

	return response, nil
}

func reqToModel(req request.CreatePostRequest) models.Post {
	return models.Post{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		Images:  req.Images,
	}
}
