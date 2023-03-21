package controller

import (
	"file-handler/internal/defines"
	"file-handler/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type FilesController interface {
	PostFile(ctx *gin.Context)
}

type filesController struct {
	svc service.FilesService
}

func NewFilesController(svc service.FilesService) FilesController {
	return &filesController{
		svc: svc,
	}
}

func (c *filesController) PostFile(ctx *gin.Context) {
	file, err := ctx.FormFile(defines.FileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid file"})
		return
	}
	if filepath.Ext(file.Filename) != defines.FileExt {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid type file"})
		return
	}
	clientId := ctx.GetHeader("client-id")
	if clientId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "missing client-id"})
		return
	}
	err = c.svc.PostFile(file, ctx, clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(201)
}
