package binder

type CreatePost struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	AuthorID   string `json:"author_id"`
}

type UpdatePost struct {
	ID         string `param:"id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	AuthorID   string `json:"author_id"`
}