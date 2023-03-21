package repository

import (
	"file-handler/internal/defines"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"mime/multipart"
	"os"
)

type FilesRepository interface {
	PostFile(string, *multipart.FileHeader, *gin.Context) error
}

type filesRepository struct {
	PathFile string
	FileExt  string
}

func NewFilesRepository() FilesRepository {
	return &filesRepository{
		PathFile: os.Getenv(defines.EnvPathSaveFile),
		FileExt:  defines.FileExt,
	}
}

func (r *filesRepository) PostFile(nameFile string, file *multipart.FileHeader, ctx *gin.Context) error {
	err := ctx.SaveUploadedFile(file, r.PathFile+nameFile+r.FileExt)
	if err != nil {
		return err
	}
	return nil
}
