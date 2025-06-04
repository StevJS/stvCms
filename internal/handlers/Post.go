package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stvCms/internal/rest/request"
)

func CreatePost(ctx *gin.Context) {
	postRequest := request.CreatePostRequest{}

	if err := ctx.ShouldBindJSON(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
