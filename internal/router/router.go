package router

import (
	"context"
	"file-handler/internal/controller"
	"file-handler/internal/defines"
	"file-handler/internal/repository"
	"file-handler/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(defines.EnvRedisHost),
	})
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}
	// Repositories init
	repo := repository.NewFilesRepository(redisClient)

	// Services init
	svc := service.NewFilesService(repo)

	// Controllers init
	ctrl := controller.NewFilesController(svc)

	// Endpoints
	r.POST(defines.EndpointPostFile, ctrl.PostFile)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))

}
