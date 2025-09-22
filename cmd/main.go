package main

import (
	"log"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/scheduler"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/providers"
	"github.com/Caknoooo/go-gin-clean-starter/script"
	"github.com/samber/do"
	"gorm.io/gorm"

	// "github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

func args(injector *do.Injector) bool {
	if len(os.Args) > 1 {
		flag := script.Commands(injector)
		return flag
	}

	return true
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	// myFigure := figure.NewColorFigure("Hafaro", "", "green", true)
	// myFigure.Print()

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
	var (
		injector = do.New()
	)

	providers.RegisterDependencies(injector)

	if !args(injector) {
		return
	}

	server := gin.Default()
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	scheduler.Start(db)

	// Atur trusted proxies, misal hanya localhost dan IP proxy tertentu
	server.SetTrustedProxies([]string{"127.0.0.1"})

	// Enable CORS middleware
	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	user.RegisterRoutes(server, injector)
	auth.RegisterRoutes(server, injector)
	tenant.RegisterRoutes(server, injector)
	product.RegisterRoutes(server, injector)

	run(server)
}
