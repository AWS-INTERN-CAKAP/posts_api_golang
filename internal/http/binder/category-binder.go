package binder

import "github.com/google/uuid"

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type UpdateCategory struct {
	ID        string `param:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type GetByID struct {
	ID string `param:"id" validate:"required"`
}