package repository

import (
	"gorm.io/gorm"
	"stvCms/internal/config"
	"stvCms/internal/models"
)

type IPostRepository interface {
	CreatePost(post models.Post) (string, error)
	GetPosts() ([]models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostGormRepository() *postRepository {
	return &postRepository{
		db: config.Init(),
	}
}

func (pr *postRepository) CreatePost(post models.Post) (string, error) {
	err := pr.db.Create(&post).Error
	if err != nil {
		return "No se pudo crear el post", err
	}
	return "Post creado", nil
}

func (pr *postRepository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := pr.db.Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}
