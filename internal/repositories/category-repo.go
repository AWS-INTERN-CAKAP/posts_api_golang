package repositories

import (
	"github.com/Davidafdal/e-commerce/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uuid.UUID) (*entity.Category, error)
	UpdateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(id uuid.UUID) error
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *entity.Category) (*entity.Category, error) {
	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id uuid.UUID) (*entity.Category, error) {
	var category entity.Category
	err := r.DB.Preload("Posts").First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	if err := r.DB.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategory(id uuid.UUID) error {
	return r.DB.Delete(&entity.Category{}, "id = ?", id).Error
}