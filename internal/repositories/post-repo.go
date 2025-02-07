package repositories

import (
	"github.com/Davidafdal/e-commerce/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	GetAllPosts() ([]entity.Post, error)
	GetPostByID(id uuid.UUID) (*entity.Post, error)
	UpdatePost(post *entity.Post) (*entity.Post, error)
	DeletePost(id uuid.UUID) error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// CreatePost - Creates a new post and returns it
func (r *postRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	if err := r.DB.Create(post).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Category").Preload("Author").First(post, "id = ?", post.ID).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// GetAllPosts - Retrieves all posts
func (r *postRepository) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post
	err := r.DB.Preload("Category").Preload("Author").Find(&posts).Error
	return posts, err
}

// GetPostByID - Retrieves a post by ID
func (r *postRepository) GetPostByID(id uuid.UUID) (*entity.Post, error) {
	var post entity.Post
	err := r.DB.Preload("Category").Preload("Author").First(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost - Updates a post and returns the updated post
func (r *postRepository) UpdatePost(post *entity.Post) (*entity.Post, error) {
	if err := r.DB.Save(post).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Category").Preload("Author").First(post, "id = ?", post.ID).Error; err != nil {
		return nil, err
	}
	
	return post, nil
}

// DeletePost - Deletes a post by ID
func (r *postRepository) DeletePost(id uuid.UUID) error {
	return r.DB.Delete(&entity.Post{}, "id = ?", id).Error
}