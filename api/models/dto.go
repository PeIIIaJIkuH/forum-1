package models

type InputUserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InputUserSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type InputPost struct {
	ID         int64    `json:"id"`
	AuthorID   int64    `json:"authorID"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

type InputComment struct {
	ID      int64  `json:"id"`
	PostID  int64  `json:"post_id"`
	Content string `json:"content"`
}

type InputFindComment struct {
	Option string `json:"option"` // user or post
	PostID int64  `json:"post_id"`
	UserID int64  `json:"user_id"`
}

type InputRate struct {
	ID       int64 `json:"id"`       // postID
	Reaction int   `json:"reaction"` // 1 or -1
}

type InputFilterPost struct {
	Option     string   `json:"option"`   // categories or date or rating or author
	AuthorID   int64    `json:"authorID"` // getAllPosts created by AuthorID
	Date       string   `json:"date"`     // ASC or DESC
	Rating     string   `json:"rating"`   // ASC or DESC
	Categories []string `json:"categories"`
	UserRating string   `json:"userRating"` // upvoted or downvoted
}
