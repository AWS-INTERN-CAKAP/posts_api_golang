package services

import (
	"errors"

	"github.com/Davidafdal/e-commerce/internal/dto"
	"github.com/Davidafdal/e-commerce/internal/entity"
	"github.com/Davidafdal/e-commerce/internal/http/binder"
	"github.com/Davidafdal/e-commerce/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrorCategoryNotFound = errors.New("data category tidak ditemukan")

type CategoryService interface {
	GetCategories() (dto.GetCategories, error)
	GetCategoryByID(id uuid.UUID) (dto.GetCategory, error)
	CreateCategory(input *binder.CreateCategory) (dto.CreateCategory, error)
	UpdateCategory(input *binder.UpdateCategory) (dto.UpdateCategory, error)
	DeleteCategory(id uuid.UUID) error
}


type categoryService struct {
	categoryRepo repositories.CategoryRepository	
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}




func (s *categoryService) GetCategories() (dto.GetCategories, error) {

	categories, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return dto.GetCategories{
			Categories: []dto.Category{},
		}, ErrorServerError
	}

	categoriesData := make([]dto.Category, 0)
	for _, category := range categories {
		categoriesData = append(categoriesData, dto.Category{
			ID:   category.ID.String(),
			Name: category.Name,
			Description: category.Description,
		})
	}

	return dto.GetCategories{Categories: categoriesData}, nil
}


func (s *categoryService) GetCategoryByID(id uuid.UUID) (dto.GetCategory, error) {

	category, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.GetCategory{
				Category: dto.CategoryData{},
			}, ErrorCategoryNotFound
		}
		return dto.GetCategory{}, ErrorServerError
	}

	posts := make([]dto.Post, 0)

	for _, post := range category.Posts {
		posts = append(posts, dto.Post{
			ID:      post.ID.String(),
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return dto.GetCategory{Category: dto.CategoryData{
		ID:   category.ID.String(),
		Name: category.Name,
		Description: category.Description,
		Posts: posts,
	}}, nil
}

func (s *categoryService) CreateCategory(input *binder.CreateCategory) (dto.CreateCategory, error) {
	categoryInput := &entity.Category{
		Name: input.Name,
		CreatedBy: input.CreatedBy,
		Description: input.Description,
	}

	newCategory, err := s.categoryRepo.CreateCategory(categoryInput)
	if err != nil {
		return dto.CreateCategory{}, ErrorServerError
    }

	return dto.CreateCategory{
		Category: dto.Category{
			ID:   newCategory.ID.String(),
			Name: newCategory.Name,
			Description: newCategory.Description,
		},
	}, nil
}


func (s *categoryService) UpdateCategory(input *binder.UpdateCategory) (dto.UpdateCategory, error) {
	categoryInput := &entity.Category{
		ID:   uuid.MustParse(input.ID),
		Name: input.Name,
		CreatedBy: input.CreatedBy,
		Description: input.Description,
	}

	updatedCategory, err := s.categoryRepo.UpdateCategory(categoryInput)
	if err != nil {
		return dto.UpdateCategory{}, ErrorServerError
	}

	return dto.UpdateCategory{
		Category: dto.Category{
			ID:   updatedCategory.ID.String(),
			Name: updatedCategory.Name,
			Description: updatedCategory.Description,
		},
	}, nil
}

func (s *categoryService) DeleteCategory(id uuid.UUID) error {
	return s.categoryRepo.DeleteCategory(id)
}




