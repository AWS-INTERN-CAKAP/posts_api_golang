package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    
	Description string   
	CreatedBy   uuid.UUID
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Posts       []Post    `gorm:"foreignKey:CategoryID"`
	Creator User `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`
}

func (Category) TableName() string {
	return "categories"
}
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

