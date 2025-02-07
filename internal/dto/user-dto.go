package dto

type UserData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserWithPostAndCategory struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Posts      []Post     `json:"posts"`
	Categories []Category `json:"categories"`
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

type RegisterResponse struct {
	User UserData `json:"user"`
}

type GetUsers struct {
	Users []UserData `json:"users"`
}

type GetUser struct {
	User UserWithPostAndCategory `json:"user"`
}
