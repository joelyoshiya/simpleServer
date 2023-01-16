package dataStructs

import "time"

// Client-side data structures

type Body struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       Authors `json:"data"`
}

type Author struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	About           string    `json:"about"`
	Submitted       int       `json:"submitted"`
	UpdatedAt       time.Time `json:"updated_at"`
	SubmissionCount int       `json:"submission_count"`
	CommentCount    int       `json:"comment_count"`
	CreatedAt       int       `json:"created_at"`
}

type Authors []Author

// Server-side data structures

type User struct {
	UserID  int    `json:"userid"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Users []User
