package main

import (
	"buildgram/handlers"
	"buildgram/middleware"
	"buildgram/store"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise the in-memory store.
	s := store.New()

	// Build handler structs that depend on the store.
	userHandler := &handlers.UserHandler{Store: s}
	postHandler := &handlers.PostHandler{Store: s}
	commentHandler := &handlers.CommentHandler{Store: s}

	// Create the Gin router without the default logger/recovery so our custom
	// logger is the only one printing to stdout.
	router := gin.New()
	router.Use(gin.Recovery()) // still recover from panics
	router.Use(middleware.RequestLogger())

	// All routes live under /api/v1 as required.
	v1 := router.Group("/api/v1")
	{
		// User routes
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUserByID)

		// Post routes
		v1.POST("/posts", postHandler.CreatePost)
		v1.GET("/posts", postHandler.GetAllPosts)
		v1.GET("/posts/:id", postHandler.GetPostByID)

		// Engagement routes
		v1.POST("/posts/:id/like", postHandler.LikePost)
		v1.POST("/posts/:id/comments", commentHandler.AddComment)
	}

	// Start the server on port 8080.
	router.Run(":8080")
}
