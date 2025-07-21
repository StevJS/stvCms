package services

import (
	"errors"
	"stvCms/internal/models"
	"stvCms/internal/repository"
	"stvCms/internal/rest/request"
	"testing"
)

type MockPostRepository struct {
	CreatePostFunc     func(post models.Post) (string, error)
	GetPostsFunc       func() ([]models.Post, error)
	UpdatePostFunc     func(id uint, post models.Post) (string, error)
	GetPostByIdFunc    func(id uint) (models.Post, error)
	DeletePostByIdFunc func(id int) bool
}

var _ repository.IPostRepository = &MockPostRepository{}

func (m *MockPostRepository) CreatePost(post models.Post) (string, error) {
	if m.CreatePostFunc != nil {
		return m.CreatePostFunc(post)
	}
	return "", errors.New("CreatePostFunc no implementado")
}

func (m *MockPostRepository) GetPosts() ([]models.Post, error) {
	if m.GetPostsFunc != nil {
		return m.GetPostsFunc()
	}
	return nil, errors.New("GetPostsFunc no implementado")
}

func (m *MockPostRepository) UpdatePost(id uint, post models.Post) (string, error) {
	if m.UpdatePostFunc != nil {
		return m.UpdatePostFunc(id, post)
	}
	return "", errors.New("UpdatePostFunc no implementado")
}

func (m *MockPostRepository) GetPostById(id uint) (models.Post, error) {
	if m.GetPostByIdFunc != nil {
		return m.GetPostByIdFunc(id)
	}
	return models.Post{}, errors.New("GetPostByIdFunc no implementado")
}

func (m *MockPostRepository) DeletePostById(id int) bool {
	if m.DeletePostByIdFunc != nil {
		return m.DeletePostByIdFunc(id)
	}
	return false
}

func TestPostService_CreatePost(t *testing.T) {
	mockRepo := &MockPostRepository{}
	service := &postService{repository: mockRepo}

	createPostRequest := request.CreatePostRequest{
		Title:   "Test Post",
		Content: "Contenido del post de prueba.",
		Author:  "Tester",
	}

	expectedResponse := "Post creado"
	mockRepo.CreatePostFunc = func(post models.Post) (string, error) {
		if post.Title != createPostRequest.Title {
			t.Errorf("Título esperado '%s', obtenido '%s'", createPostRequest.Title, post.Title)
		}
		if post.Content != createPostRequest.Content {
			t.Errorf("Contenido esperado '%s', obtenido '%s'", createPostRequest.Content, post.Content)
		}
		if post.Author != createPostRequest.Author {
			t.Errorf("Autor esperado '%s', obtenido '%s'", createPostRequest.Author, post.Author)
		}
		return expectedResponse, nil
	}

	response, err := service.CreatePost(createPostRequest)

	if err != nil {
		t.Fatalf("CreatePost() devolvió un error inesperado: %v", err)
	}

	if response != expectedResponse {
		t.Errorf("Respuesta esperada '%s', obtenida '%s'", expectedResponse, response)
	}
}
