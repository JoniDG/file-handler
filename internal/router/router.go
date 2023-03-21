package router

import (
	"file-handler/internal/controller"
	"file-handler/internal/defines"
	"file-handler/internal/repository"
	"file-handler/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init

	// Repositories init
	repo := repository.NewFilesRepository()

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
