package controller

import (
	"net/http"
	"strconv"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/gin-gonic/gin"
)

// CommentController is responsible for handling the post input and output of the application
type PostControler struct {
	postService usecase.PostService
}

// NewPostController creates a new PostController instance
func NewPostController(postService usecase.PostService) *PostControler {
	return &PostControler{postService: postService}
}

// CreatePost creates a new Post
func (c *PostControler) CreatePost(ctx *gin.Context) {
	var input entity.Post
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := c.postService.CreatePost(input.Creator, input.Subreddit, input.Title, input.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data saved successfully": post})
}

// GetPost fetches post based on the postID
func (c *PostControler) GetPost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	post, err := c.postService.GetPost(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

// ListPosts fetches the application posts
func (c *PostControler) ListPosts(ctx *gin.Context) {
	posts, err := c.postService.ListPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}
