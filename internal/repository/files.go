package repository

import (
	"context"
	"file-handler/internal/defines"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"mime/multipart"
	"os"
)

type FilesRepository interface {
	PostFile(nameFile string, file *multipart.FileHeader, ctx *gin.Context) error
}

type filesRepository struct {
	PathFile    string
	FileExt     string
	redisClient *redis.Client
}

func NewFilesRepository(redisClient *redis.Client) FilesRepository {
	return &filesRepository{
		PathFile:    os.Getenv(defines.EnvPathSaveFile),
		FileExt:     defines.FileExt,
		redisClient: redisClient,
	}
}

func (r *filesRepository) PostFile(nameFile string, file *multipart.FileHeader, ctxGin *gin.Context) error {
	err := ctxGin.SaveUploadedFile(file, r.PathFile+nameFile+r.FileExt)
	if err != nil {
		return err
	}
	ctxRedis := context.Background()
	return r.redisClient.RPush(ctxRedis, defines.QueueUploadFile, r.PathFile+nameFile+r.FileExt).Err()
}
