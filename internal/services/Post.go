package services

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"stvCms/internal/models"
	"stvCms/internal/repository"
	"stvCms/internal/rest/request"
	"stvCms/internal/rest/response"
)

type IPostService interface {
	CreatePost(req request.CreatePostRequest) (string, error)
	GetPosts() ([]response.PostResponse, error)
	GetPostById(id string) (response.PostResponse, error)
	UpdatePost(req request.UpdatePostRequest) (string, error)
	DeletePostById(id string) (string, error)
	InsertCodeContentInPost(content request.CodeContent) (string, error)
	//GetCodeContent()
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

	modelPost, err := ps.repository.CreatePost(post)
	if err != nil {
		return "No se pudo crear el post", err
	}

	return modelPost, nil
}

func (ps *postService) GetPosts() ([]response.PostResponse, error) {

	posts := []response.PostResponse{}
	modelPosts, err := ps.repository.GetPosts()
	if err != nil {
		return posts, nil
	}

	for _, post := range modelPosts {
		data := response.PostResponse{
			Id:          post.Model.ID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Title:       post.Title,
			Content:     post.Content,
			Author:      post.Author,
			CodeContent: ps.GetCodeContent(post.ID),
		}
		posts = append(posts, data)
	}

	return posts, nil
}

func (ps *postService) GetCodeContent(postID uint) []response.CodeContentResponse {
	if postID == 0 {
		return []response.CodeContentResponse{}
	}

	codeContents, err := ps.repository.GetCodeContents(postID)

	if err != nil {
		return []response.CodeContentResponse{}
	}
	var codeContentResponses []response.CodeContentResponse

	for _, codeContent := range codeContents {
		codeContentResponse := response.CodeContentResponse{
			Code:     codeContent.Code,
			Language: codeContent.Language,
		}

		codeContentResponses = append(codeContentResponses, codeContentResponse)
	}

	return codeContentResponses
}

func (ps *postService) GetPostById(id string) (response.PostResponse, error) {
	postId, _ := strconv.Atoi(id)

	post, err := ps.repository.GetPostById(uint(postId))

	postResponse := response.PostResponse{
		Id:        post.Model.ID,
		CreatedAt: post.CreatedAt,
		Title:     post.Title,
		Content:   post.Content,
		Author:    post.Author,
	}

	if err != nil {
		return postResponse, err
	}

	return postResponse, nil
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

func (ps *postService) DeletePostById(id string) (string, error) {
	postId, _ := strconv.Atoi(id)

	ok := ps.repository.DeletePostById(postId)

	if !ok {
		return "", fmt.Errorf("Error al borrar el post")
	}

	return "Post borrado", nil
}

func reqToModel(req request.CreatePostRequest) models.Post {
	return models.Post{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		//Images:  req.Images,
	}
}

func (ps *postService) InsertCodeContentInPost(request request.CodeContent) (string, error) {

	model := models.CodeContent{
		Code:     request.Content.Code,
		Language: request.Content.Language,
		PostID:   uint(request.PostID),
	}
	err := ps.repository.SaveCodeContentInPost(model)
	if err != nil {
		return "", err
	}

	return "Code content asociado correctamente", nil
}
