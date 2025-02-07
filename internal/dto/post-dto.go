package dto

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostData struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category Category `json:"category"`
	Author   UserData `json:"author"`
}

type GetPosts struct {
	Posts []PostData `json:"posts"`
}

type GetPost struct {
	Post PostData `json:"post"`
}

type CreatePost struct {
	Post PostData `json:"post"`
}

type UpdatePost struct {
	Post PostData `json:"post"`
}