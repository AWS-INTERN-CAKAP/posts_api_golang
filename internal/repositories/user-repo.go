package repositories

import (
	"github.com/Davidafdal/e-commerce/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserByID(id uuid.UUID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) (bool, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}


func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}


func (r *userRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Preload("Categories").Preload("Posts").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}


func (r *userRepository) DeleteUser(user *entity.User) (bool, error) {
	if err := r.DB.Delete(user).Error; err != nil {
		return false, err
	}

	return true, nil
}