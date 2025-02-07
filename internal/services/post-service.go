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

var ErrorPostNotFound = errors.New("data post tidak ditemukan")

type PostService interface {
	GetPosts() (dto.GetPosts, error)
	GetPostByID(id uuid.UUID) (dto.GetPost, error)
	CreatePost(data *binder.CreatePost) (dto.CreatePost, error)
	UpdatePost(input *binder.UpdatePost) (dto.UpdatePost, error)
	DeletePost(id uuid.UUID) error
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostService {
	return &postService{postRepo}
}


func (s *postService) GetPosts() (dto.GetPosts, error) {	
	posts, err := s.postRepo.GetAllPosts()

	if err != nil {
		return dto.GetPosts{
			Posts: []dto.PostData{},
		}, ErrorServerError
	}

	postsData := make([]dto.PostData, 0)
	for _, post := range posts {
		postsData = append(postsData, dto.PostData{
			ID:         post.ID.String(),
			Title:      post.Title,
			Content:    post.Content,
			Category: dto.Category{
				ID:   post.Category.ID.String(),
				Name: post.Category.Name,
			},
			Author: dto.UserData{
				ID:   post.Author.ID.String(),
				Name: post.Author.Name,
			},
		})
	}

	return dto.GetPosts{Posts: postsData}, nil
}

func (s *postService) GetPostByID(id uuid.UUID) (dto.GetPost, error) {

	post, err := s.postRepo.GetPostByID(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return dto.GetPost{
				Post: dto.PostData{},
			}, ErrorPostNotFound
		}

		return dto.GetPost{
			Post: dto.PostData{},
		}, ErrorServerError
	}

	return dto.GetPost{Post: dto.PostData{
		ID:         post.ID.String(),
		Title:      post.Title,
		Content:    post.Content,
		Category: dto.Category{
			ID:   post.Category.ID.String(),
			Name: post.Category.Name,
		},
		Author: dto.UserData{
			ID:   post.Author.ID.String(),
			Name: post.Author.Name,
		},
	}}, nil
}

func (s *postService) CreatePost(input *binder.CreatePost) (dto.CreatePost, error) {
	post, err := s.postRepo.CreatePost(&entity.Post{
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: uuid.MustParse(input.CategoryID),
		AuthorID:   uuid.MustParse(input.AuthorID),
	})

	if err != nil {
		return dto.CreatePost{
			Post: dto.PostData{},
		}, ErrorServerError
	}


	return dto.CreatePost{Post: dto.PostData{
		ID:         post.ID.String(),
		Title:      post.Title,
		Content:    post.Content,
		Category: dto.Category{
			ID:   post.Category.ID.String(),
			Name: post.Category.Name,
		},
		Author: dto.UserData{
			ID:   post.Author.ID.String(),
			Name: post.Author.Name,
		},
	}}, nil
}

func (s *postService) UpdatePost(input *binder.UpdatePost) (dto.UpdatePost, error) {
	post, err := s.postRepo.UpdatePost(&entity.Post{
		ID:         uuid.MustParse(input.ID),
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: uuid.MustParse(input.CategoryID),
		AuthorID:   uuid.MustParse(input.AuthorID),
	})

	if err != nil {
		return dto.UpdatePost{
			Post: dto.PostData{},
		}, ErrorServerError
	}

	return dto.UpdatePost{Post: dto.PostData{
		ID:         post.ID.String(),
		Title:      post.Title,
		Content:    post.Content,
		Category: dto.Category{
			ID:   post.Category.ID.String(),
			Name: post.Category.Name,
		},
		Author: dto.UserData{
			ID:   post.Author.ID.String(),
			Name: post.Author.Name,
		},
	}}, nil
}

func (s *postService) DeletePost(id uuid.UUID) error {
	return s.postRepo.DeletePost(id)
}


