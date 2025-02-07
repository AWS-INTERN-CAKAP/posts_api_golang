package dto

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CategoryData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"posts"`
}

type GetCategories struct {
	Categories []Category `json:"categories"`
}

type GetCategory struct {
	Category CategoryData `json:"category"`
}

type CreateCategory struct {
	Category Category `json:"category"`
}

type UpdateCategory struct {
	Category Category `json:"category"`
}
