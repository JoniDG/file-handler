package service

import (
	"file-handler/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
)

type FilesService interface {
	PostFile(*multipart.FileHeader, *gin.Context, string) (*string, error)
}

type filesService struct {
	repo repository.FilesRepository
}

func NewFilesService(repo repository.FilesRepository) FilesService {
	return &filesService{
		repo: repo,
	}
}

func (r *filesService) PostFile(file *multipart.FileHeader, ctx *gin.Context, clientId string) (*string, error) {
	id := uuid.New()
	nameFile := fmt.Sprintf("%s-%s", clientId, id)
	ok, err := r.repo.PostFile(nameFile, file, ctx)
	if err != nil {
		return nil, err
	}
	return ok, nil
}
