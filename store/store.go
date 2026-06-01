// Package store provides a thread-safe in-memory data store for BuildGram.
package store

import (
	"buildgram/models"
	"sync"
)

// Store is the in-memory database holding users, posts, and comments.
type Store struct {
	mu       sync.RWMutex
	users    map[int]models.User
	posts    map[int]models.Post
	comments map[int]models.Comment

	nextUserID    int
	nextPostID    int
	nextCommentID int
}

// New creates and returns a new, empty Store.
func New() *Store {
	return &Store{
		users:         make(map[int]models.User),
		posts:         make(map[int]models.Post),
		comments:      make(map[int]models.Comment),
		nextUserID:    1,
		nextPostID:    1,
		nextCommentID: 1,
	}
}

// ---------- User operations ----------

// CreateUser adds a new user to the store and returns it with its assigned ID.
func (s *Store) CreateUser(u models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	u.ID = s.nextUserID
	s.nextUserID++
	s.users[u.ID] = u
	return u
}

// GetUserByID returns the user with the given ID and a boolean indicating
// whether the user was found.
func (s *Store) GetUserByID(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	return u, ok
}

// UserExists returns true if a user with the given ID exists.
func (s *Store) UserExists(id int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.users[id]
	return ok
}

// ---------- Post operations ----------

// CreatePost adds a new post to the store and returns it with its assigned ID.
func (s *Store) CreatePost(p models.Post) models.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = s.nextPostID
	s.nextPostID++
	s.posts[p.ID] = p
	return p
}

// GetAllPosts returns a slice of all posts in the store.
func (s *Store) GetAllPosts() []models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]models.Post, 0, len(s.posts))
	for _, p := range s.posts {
		posts = append(posts, p)
	}
	return posts
}

// GetPostByID returns the post with the given ID and a boolean indicating
// whether the post was found.
func (s *Store) GetPostByID(id int) (models.Post, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.posts[id]
	return p, ok
}

// LikePost increments the LikesCount on the post with the given ID.
// It returns the updated post and a boolean indicating success.
func (s *Store) LikePost(id int) (models.Post, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.posts[id]
	if !ok {
		return models.Post{}, false
	}
	p.LikesCount++
	s.posts[id] = p
	return p, true
}

// ---------- Comment operations ----------

// CreateComment adds a new comment to the store and returns it with its assigned ID.
func (s *Store) CreateComment(c models.Comment) models.Comment {
	s.mu.Lock()
	defer s.mu.Unlock()

	c.ID = s.nextCommentID
	s.nextCommentID++
	s.comments[c.ID] = c
	return c
}

// GetCommentsByPostID returns all comments associated with the given post ID.
func (s *Store) GetCommentsByPostID(postID int) []models.Comment {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Comment
	for _, c := range s.comments {
		if c.PostID == postID {
			result = append(result, c)
		}
	}
	return result
}
