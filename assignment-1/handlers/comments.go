package handlers

import (
	"buildgram/models"
	"buildgram/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CommentHandler holds the store dependency for comment-related handlers.
type CommentHandler struct {
	Store *store.Store
}

// addCommentRequest is the expected request body for POST /api/v1/posts/:id/comments.
type addCommentRequest struct {
	UserID int    `json:"userID" binding:"required"`
	Text   string `json:"text"   binding:"required"`
}

// AddComment handles POST /api/v1/posts/:id/comments.
// It validates the body, checks that both the post and the user exist, then
// creates and returns the new comment with HTTP 201.
func (h *CommentHandler) AddComment(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		badRequest(c, "id must be a valid integer")
		return
	}

	// Verify the post exists before binding the body to avoid wasting work.
	if _, ok := h.Store.GetPostByID(postID); !ok {
		notFound(c, "post not found")
		return
	}

	var req addCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, "invalid request body: userID and text are required")
		return
	}

	if !h.Store.UserExists(req.UserID) {
		notFound(c, "user not found: no user exists with the provided userID")
		return
	}

	comment := h.Store.CreateComment(models.Comment{
		PostID:    postID,
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now().UTC(),
	})

	successResponse(c, http.StatusCreated, comment)
}
