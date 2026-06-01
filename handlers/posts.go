package handlers

import (
	"buildgram/models"
	"buildgram/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PostHandler holds the store dependency for post-related handlers.
type PostHandler struct {
	Store *store.Store
}

// createPostRequest is the expected request body for POST /api/v1/posts.
type createPostRequest struct {
	UserID   int     `json:"userID"   binding:"required"`
	ImageURL string  `json:"imageURL" binding:"required"`
	Caption  *string `json:"caption"` // optional
}

// CreatePost handles POST /api/v1/posts.
// It validates the body, ensures the referenced user exists, and creates a post.
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, "invalid request body: userID and imageURL are required")
		return
	}

	if !h.Store.UserExists(req.UserID) {
		notFound(c, "user not found: no user exists with the provided userID")
		return
	}

	post := h.Store.CreatePost(models.Post{
		UserID:     req.UserID,
		ImageURL:   req.ImageURL,
		Caption:    req.Caption,
		Timestamp:  time.Now().UTC(),
		LikesCount: 0,
	})

	successResponse(c, http.StatusCreated, post)
}

// GetAllPosts handles GET /api/v1/posts.
// It returns every post in the store as a global feed.
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts := h.Store.GetAllPosts()

	// Return an empty array instead of null when there are no posts.
	if posts == nil {
		posts = []models.Post{}
	}

	successResponse(c, http.StatusOK, posts)
}

// GetPostByID handles GET /api/v1/posts/:id.
// It returns the post together with all of its associated comments.
func (h *PostHandler) GetPostByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequest(c, "id must be a valid integer")
		return
	}

	post, ok := h.Store.GetPostByID(id)
	if !ok {
		notFound(c, "post not found")
		return
	}

	comments := h.Store.GetCommentsByPostID(id)
	if comments == nil {
		comments = []models.Comment{}
	}

	successResponse(c, http.StatusOK, gin.H{
		"post":     post,
		"comments": comments,
	})
}

// LikePost handles POST /api/v1/posts/:id/like.
// It increments the like count on the specified post.
func (h *PostHandler) LikePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequest(c, "id must be a valid integer")
		return
	}

	post, ok := h.Store.LikePost(id)
	if !ok {
		notFound(c, "post not found")
		return
	}

	successResponse(c, http.StatusOK, gin.H{
		"id":         post.ID,
		"likesCount": post.LikesCount,
	})
}
