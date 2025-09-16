package product

import (
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product/controller"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	productController := do.MustInvoke[controller.ProductController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	productRoutes := server.Group("/api/product")
	{
		productRoutes.POST("", middlewares.Authenticate(jwtService), productController.Create)
		productRoutes.GET("", middlewares.Authenticate(jwtService), productController.GetAll)
		productRoutes.GET("/:id", middlewares.Authenticate(jwtService), productController.GetById)
		productRoutes.PUT("/:id", middlewares.Authenticate(jwtService), productController.Update)
		productRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), productController.Delete)
	}
}
