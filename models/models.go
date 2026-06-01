package models

import "time"

// User represents a registered user in BuildGram.
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio"` // nullable — a user may not have a bio
}

// Post represents an image post made by a user.
type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userID"`
	ImageURL   string    `json:"imageURL"`
	Caption    *string   `json:"caption"` // nullable — a post may have no caption
	Timestamp  time.Time `json:"timestamp"`
	LikesCount int       `json:"likesCount"`
}

// Comment represents a comment made by a user on a post.
type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postID"`
	UserID    int       `json:"userID"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
