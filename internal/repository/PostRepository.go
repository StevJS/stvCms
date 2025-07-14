package repository

import (
	"gorm.io/gorm"
	"stvCms/internal/config"
	"stvCms/internal/models"
)

type IPostRepository interface {
	CreatePost(post models.Post) (string, error)
	GetPosts() ([]models.Post, error)
	UpdatePost(id uint, post models.Post) (string, error)
	GetPostById(id uint) (models.Post, error)
	DeletePostById(id int) bool
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

func (pr *postRepository) UpdatePost(id uint, post models.Post) (string, error) {
	result := pr.db.Model(&models.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		return "No se pudo actualizar el post", result.Error
	}

	if result.RowsAffected == 0 {
		return "El post no fue encontrado o no habían datos para actualizar", nil
	}

	return "Post actualizado", nil
}

func (pr *postRepository) GetPostById(id uint) (models.Post, error) {
	var post models.Post
	err := pr.db.First(&post, id).Error
	return post, err
}

func (pr *postRepository) DeletePostById(id int) bool {
	ok := pr.db.Delete(&models.Post{}, id).RowsAffected > 0
	return ok
}
