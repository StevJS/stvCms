package services

import (
	"gorm.io/gorm"
	"stvCms/internal/models"
	"stvCms/internal/repository"
	"stvCms/internal/rest/request"
	"stvCms/internal/rest/response"
)

type IPostService interface {
	CreatePost(req request.CreatePostRequest) (string, error)
	GetPosts() ([]response.PostResponse, error)
	UpdatePost(req request.UpdatePostRequest) (string, error)
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
			Id:      post.Model.ID,
			Title:   post.Title,
			Content: post.Content,
			Author:  post.Author,
		}
		posts = append(posts, data)
	}

	return posts, nil
}

func (ps *postService) UpdatePost(req request.UpdatePostRequest) (string, error) {
	postModel, err := ps.repository.GetPostById(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Post no encontrado con ese ID", err
		}
		return "Error al buscar el post", err
	}

	// mapping req to model
	postModel.Title = req.Title
	postModel.Content = req.Content
	postModel.Author = req.Author

	postUpdated, err := ps.repository.UpdatePost(req.Id, postModel)
	if err != nil {
		return "", err
	}

	return postUpdated, nil
}

func reqToModel(req request.CreatePostRequest) models.Post {
	return models.Post{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		//Images:  req.Images,
	}
}
