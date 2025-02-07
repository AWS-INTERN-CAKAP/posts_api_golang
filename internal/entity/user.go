package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string     `gorm:"type:varchar(100);not null"`
	Email      string     `gorm:"type:varchar(100);unique;not null"`
	Password   string     `gorm:"type:text;not null"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	Posts      []Post     `gorm:"foreignKey:AuthorID"`
	Categories []Category `gorm:"foreignKey:CreatedBy"`
}


func (User) TableName() string {
    return "users" 
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}