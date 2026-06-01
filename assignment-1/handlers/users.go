package handlers

import (
	"buildgram/models"
	"buildgram/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler holds the store dependency for user-related handlers.
type UserHandler struct {
	Store *store.Store
}

// createUserRequest is the expected request body for POST /api/v1/users.
type createUserRequest struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email"    binding:"required,email"`
	Bio      *string `json:"bio"` // optional — pointer so absence is distinguishable from empty string
}

// CreateUser handles POST /api/v1/users.
// It creates a new user and returns the created user with HTTP 201.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badRequest(c, "invalid request body: username is required, and email must be a valid email address")
		return
	}

	user, err := h.Store.CreateUser(models.User{
		Username: req.Username,
		Email:    req.Email,
		Bio:      req.Bio,
	})
	if err != nil {
		if err == store.ErrUsernameTaken || err == store.ErrEmailTaken {
			errorResponse(c, http.StatusConflict, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	successResponse(c, http.StatusCreated, user)
}

// GetUserByID handles GET /api/v1/users/:id.
// It returns the profile of the user with the given ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		badRequest(c, "id must be a valid integer")
		return
	}

	user, ok := h.Store.GetUserByID(id)
	if !ok {
		notFound(c, "user not found")
		return
	}

	successResponse(c, http.StatusOK, user)
}
