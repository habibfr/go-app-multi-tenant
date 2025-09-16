package providers

import (
	"github.com/Caknoooo/go-gin-clean-starter/config"
	authRepo "github.com/Caknoooo/go-gin-clean-starter/modules/auth/repository"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	productController "github.com/Caknoooo/go-gin-clean-starter/modules/product/controller"
	productRepo "github.com/Caknoooo/go-gin-clean-starter/modules/product/repository"
	productService "github.com/Caknoooo/go-gin-clean-starter/modules/product/service"
	tenantController "github.com/Caknoooo/go-gin-clean-starter/modules/tenant/controller"
	tenantRepo "github.com/Caknoooo/go-gin-clean-starter/modules/tenant/repository"
	tenantService "github.com/Caknoooo/go-gin-clean-starter/modules/tenant/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/repository"
	userService "github.com/Caknoooo/go-gin-clean-starter/modules/user/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	// Initialize
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	userRepository := repository.NewUserRepository(db)
	tenantRepository := tenantRepo.NewTenantRepository(db)
	refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)
	productRepository := productRepo.NewProductRepository(db)

	// Service
	userService := userService.NewUserService(userRepository, refreshTokenRepository, jwtService, db)
	tenantService := tenantService.NewTenantService(tenantRepository, db)
	productService := productService.NewProductService(productRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.UserController, error) {
			return controller.NewUserController(i, userService), nil
		},
	)
	do.Provide(
		injector, func(i *do.Injector) (tenantController.TenantController, error) {
			return tenantController.NewTenantController(i, tenantService), nil
		},
	)
	do.Provide(
		injector, func(i *do.Injector) (productController.ProductController, error) {
			return productController.NewProductController(i, productService), nil
		},
	)
}
