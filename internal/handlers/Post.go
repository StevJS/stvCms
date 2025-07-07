package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stvCms/internal/rest/request"
	"stvCms/internal/services"
)

type postHandler struct {
	service services.IPostService
}

func NewPostHandler() *postHandler {
	return &postHandler{
		service: services.NewPostService(),
	}
}

func (h *postHandler) CreatePost(ctx *gin.Context) {
	postRequest := request.CreatePostRequest{}

	if err := ctx.ShouldBindJSON(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response, err := h.service.CreatePost(postRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})

}

func (h *postHandler) GetPosts(ctx *gin.Context) {

}
