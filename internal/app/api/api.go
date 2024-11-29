package api

import (
	"http-server/docs"
	"http-server/internal/app/api/handler"
	"http-server/internal/app/api/routes"
	"http-server/internal/app/api/services"
	"http-server/internal/shared"
	"http-server/internal/shared/database/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type ApiServer struct {
	g          *gin.Engine
	rgw        *routes.RoutesGateWay
	sharedDeps *shared.SharedDeps
}

func NewGinServer() *gin.Engine {
	return gin.New()
}
func NewApiServer(g *gin.Engine, rgw *routes.RoutesGateWay,
	sharedDeps *shared.SharedDeps) *ApiServer {
	return &ApiServer{
		g:          g,
		rgw:        rgw,
		sharedDeps: sharedDeps,
	}
}

func (api *ApiServer) StartServer() {
	// set up all routes
	api.g.Use(api.sharedDeps.Middlewares.CreateErrorHandlingMiddleware())
	api.g.Use(gin.Logger())
	docs.SwaggerInfo.BasePath = ""
	api.g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Define a custom log writer with filtering logic
	api.rgw.SetUp()

	// midleware setup
	api.g.Run(":" + api.sharedDeps.Env.APP_PORT)
}

var ApiServerModule = fx.Options(
	shared.SharedModuleFx,
	repository.Repositories,
	handler.Handlers,
	services.Services,
	routes.RoutesGateWayModule,
	fx.Provide(NewApiServer),
	fx.Provide(NewGinServer),
)
