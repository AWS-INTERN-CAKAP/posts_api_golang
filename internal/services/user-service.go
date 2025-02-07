package services

import (
	"errors"
	"time"

	"github.com/Davidafdal/e-commerce/internal/dto"
	"github.com/Davidafdal/e-commerce/internal/entity"
	"github.com/Davidafdal/e-commerce/internal/http/binder"
	"github.com/Davidafdal/e-commerce/internal/repositories"
	"github.com/Davidafdal/e-commerce/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Login(input *binder.LoginRequst) (dto.LoginResponse, error)
	Register(input *binder.RegisterRequest) (dto.RegisterResponse, error)
	GetUsers() (dto.GetUsers, error)
	GetUserByID(id string) (dto.GetUser, error)
	Logout(tokenString string) error
}

type userService struct {
	userRepository repositories.UserRepository
	tokenUseCase   token.TokenUseCase
}

var ErrorInvalidUsers = errors.New("email/password yang anda masukkan salah")
var ErrorServerError = errors.New("kesalahan pada server")
var ErrorExitedUser = errors.New("email yang anda masukkan sudah terdaftar")
var ErrorUserNotFound = errors.New("data user tidak ditemukan")


func NewUserService(userRepository repositories.UserRepository, tokenUseCase token.TokenUseCase) UserService {
	return &userService{userRepository, tokenUseCase}
}


func (s *userService) Login(input *binder.LoginRequst) (dto.LoginResponse, error) {
	user, err := s.userRepository.GetUserByEmail(input.Email)

	if err != nil {
		return dto.LoginResponse{}, ErrorInvalidUsers
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return dto.LoginResponse{}, ErrorInvalidUsers
	}

	expiredTime := time.Now().Local().Add(30 * time.Minute)

	claims := token.JwtCustomClaims{
		ID:    user.ID.String(),
		Nama:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "post-api",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}


	token, err := s.tokenUseCase.GenerateAccessToken(claims)

	if err != nil {
		return dto.LoginResponse{}, ErrorServerError
	}


	response := dto.LoginResponse{
		Token: token,
		User: dto.UserData{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		},
	}


	return response, nil
}

func (s *userService) Register(input *binder.RegisterRequest) (dto.RegisterResponse, error) {
   exitedUsers, err := s.userRepository.GetUserByEmail(input.Email)

   if err != gorm.ErrRecordNotFound && err != nil {
		return dto.RegisterResponse{}, ErrorServerError
	}

	if exitedUsers != nil {
		return dto.RegisterResponse{}, ErrorExitedUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return dto.RegisterResponse{}, ErrorServerError
	}

	user := &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	dataUser, err := s.userRepository.CreateUser(user); 


	if err != nil {
		return dto.RegisterResponse{}, ErrorServerError
	}


	response := dto.RegisterResponse{
		User: dto.UserData{
			ID:    dataUser.ID.String(),
			Name:  dataUser.Name,
			Email: dataUser.Email,
		},
	}

	return response, nil   
}
func (s *userService) GetUsers() (dto.GetUsers, error) {
	users, err := s.userRepository.GetUsers()

	if err != nil {
		return dto.GetUsers{}, ErrorServerError
	}

	usersData := make([]dto.UserData, 0)
	for _, user := range users {
		usersData = append(usersData, dto.UserData{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return dto.GetUsers{Users: usersData}, nil
}

func (s *userService) GetUserByID(id string) (dto.GetUser, error) {
	userID := uuid.MustParse(id)

	user, err := s.userRepository.GetUserByID(userID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.GetUser{}, ErrorUserNotFound
		}
		return dto.GetUser{}, ErrorServerError
	}

	posts := make([]dto.Post, 0)

	for _, post := range user.Posts {
		posts = append(posts, dto.Post{
			ID:      post.ID.String(),
			Title:   post.Title,
			Content: post.Content,
		})
	}

	categories := make([]dto.Category, 0)

	for _, category := range user.Categories {
		categories = append(categories, dto.Category{
			ID:   category.ID.String(),
			Name: category.Name,
			Description: category.Description,
		})
	}


	return dto.GetUser{User: dto.UserWithPostAndCategory{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Posts: posts,
		Categories: categories,
	}}, nil
}

func (s *userService) Logout(tokenString string) error {
	if s.tokenUseCase == nil {
		return errors.New("tokenUseCase is not initialized")
	}
	return s.tokenUseCase.InvalidateToken(tokenString)
}


