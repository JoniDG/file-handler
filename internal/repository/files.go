package repository

import (
	"file-handler/internal/defines"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FilesRepository interface {
	PostFile(string, *multipart.FileHeader, *gin.Context) (*string, error)
}

type filesRepository struct {
	PathFile string
	FileExt  string
}

func NewFilesRepository() FilesRepository {
	return &filesRepository{
		PathFile: defines.PathSaveFile,
		FileExt:  defines.FileExt,
	}
}

func (r *filesRepository) PostFile(nameFile string, file *multipart.FileHeader, ctx *gin.Context) (*string, error) {
	err := ctx.SaveUploadedFile(file, r.PathFile+nameFile+r.FileExt)
	if err != nil {
		return nil, err
	}
	return &nameFile, nil
}
