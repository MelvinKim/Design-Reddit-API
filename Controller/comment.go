package controller

import (
	"net/http"
	"strconv"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService usecase.CommentService
}

func NewCommentController(commentService usecase.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var payload entity.Comment
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := c.commentService.CreateComment(payload.Creator, payload.Post, payload.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data saved successfully": comment})
}

func (c *CommentController) ListComments(ctx *gin.Context) {
	postIDStr := ctx.Param("userID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PostID"})
		return
	}
	userIDStr := ctx.Param("postID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	comments, err := c.commentService.ListComments(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"comments": comments})
}
