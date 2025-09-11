package tenant

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	tenantController := do.MustInvoke[controller.TenantController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	userRoutes := server.Group("/api/tenant")
	{
		userRoutes.POST("", tenantController.Create)
		userRoutes.GET("", tenantController.GetAll)
		userRoutes.GET("/:id", tenantController.GetById)
		userRoutes.PUT("/:id", middlewares.Authenticate(jwtService), tenantController.Update)
		userRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), tenantController.Delete)
	}
}
