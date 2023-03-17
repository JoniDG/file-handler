package service

import (
	"file-handler/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"time"
)

type FilesService interface {
	PostFile(*multipart.FileHeader, *gin.Context) (*string, error)
}

type filesService struct {
	repo repository.FilesRepository
}

func NewFilesService(repo repository.FilesRepository) FilesService {
	return &filesService{
		repo: repo,
	}
}

func (r *filesService) PostFile(file *multipart.FileHeader, ctx *gin.Context) (*string, error) {
	now := time.Now()
	str := now.Format("02-01-2006 15:04:05")
	mili := now.Nanosecond() / 1000000
	nameFile := fmt.Sprintf(str+":%03d", mili)
	ok, err := r.repo.PostFile(nameFile, file, ctx)
	if err != nil {
		return nil, err
	}
	return ok, nil
}
