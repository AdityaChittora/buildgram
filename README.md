# BuildGram — Assignment 1

## 1. Project Overview

BuildGram is a foundational REST API for a simplified Instagram-like service,
built as part of the *Backend Development with Go* course.

The API is written in **Go** using the **Gin** web framework and stores all data
**in-memory** (no external database). It exposes endpoints for user management,
post creation, a global feed, liking posts, and commenting — everything needed
to understand REST API design, routing, JSON handling, and validation.

---

## 2. How to Run

### Prerequisites
- **Go 1.21+** — verify with `go version`

### Install dependencies
```bash
go mod tidy
```

### Start the server
```bash
go run .
```

The server starts on **port 8080**.

```
[BuildGram] server listening on :8080
```

> **Note:** All data is stored in memory. Restarting the server resets all data.
> This is expected and intentional for this assignment.

---

## 3. Project Structure

```
buildgram/
├── main.go              # Entry point: wires store, handlers, middleware, router
├── go.mod / go.sum      # Module definition and dependency lock
├── models/
│   └── models.go        # Core data types: User, Post, Comment
├── store/
│   └── store.go         # Thread-safe in-memory data store
├── handlers/
│   ├── response.go      # Shared successResponse / errorResponse helpers
│   ├── users.go         # POST /users, GET /users/:id
│   ├── posts.go         # POST /posts, GET /posts, GET /posts/:id, POST /posts/:id/like
│   └── comments.go      # POST /posts/:id/comments
├── middleware/
│   └── logger.go        # Custom request logger middleware
└── README.md
```

---

## 4. API Reference

All endpoints are prefixed with `/api/v1`.

| Method | Path | Description | Required body fields |
|--------|------|-------------|----------------------|
| `POST` | `/api/v1/users` | Create a new user | `username`, `email` |
| `GET` | `/api/v1/users/:id` | Get a user by ID | — |
| `POST` | `/api/v1/posts` | Create a new post | `userID`, `imageURL` |
| `GET` | `/api/v1/posts` | List all posts (global feed) | — |
| `GET` | `/api/v1/posts/:id` | Get a single post with its comments | — |
| `POST` | `/api/v1/posts/:id/like` | Like a post (increments count) | — |
| `POST` | `/api/v1/posts/:id/comments` | Add a comment to a post | `userID`, `text` |

### Example: Create a User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"harshit_is_sleeping","email":"harshit@example.com","bio":"Bio of Harshit"}'
```
**Response (201 Created):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "username": "harshit_is_sleeping",
    "email": "harshit@example.com",
    "bio": "Bio of Harshit"
  }
}
```

### Example: Create a Post
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"userID":1,"imageURL":"https://example.com/photo.jpg","caption":"My first post!"}'
```
**Response (201 Created):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "userID": 1,
    "imageURL": "https://example.com/photo.jpg",
    "caption": "My first post!",
    "timestamp": "2025-06-01T10:00:00Z",
    "likesCount": 0
  }
}
```

### Example: Like a Post
```bash
curl -X POST http://localhost:8080/api/v1/posts/1/like
```
**Response (200 OK):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "likesCount": 1
  }
}
```

### Example: Add a Comment
```bash
curl -X POST http://localhost:8080/api/v1/posts/1/comments \
  -H "Content-Type: application/json" \
  -d '{"userID":2,"text":"This is stunning!"}'
```
**Response (201 Created):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "postID": 1,
    "userID": 2,
    "text": "This is stunning!",
    "timestamp": "2025-06-01T10:05:00Z"
  }
}
```

### Error Responses

All errors follow a consistent format:
```json
{
  "status": "error",
  "message": "human readable explanation"
}
```

| Scenario | Status Code |
|----------|-------------|
| Missing required field in request body | `400 Bad Request` |
| Non-integer `:id` URL parameter | `400 Bad Request` |
| User / Post not found | `404 Not Found` |
| Server panic (recovered) | `500 Internal Server Error` |

---

## 5. Middleware

A custom **Request Logger** middleware is registered globally and logs every
request to stdout:

```
[BuildGram] POST /api/v1/users | 204.75µs
[BuildGram] GET /api/v1/posts | 87.10µs
```
