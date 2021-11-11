package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/citiaps/template-go-rest/controller"
	"github.com/citiaps/template-go-rest/middleware"
	"github.com/citiaps/template-go-rest/model"
	"github.com/citiaps/template-go-rest/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/citiaps/template-go-rest/docs"
)

// @title Documentacion template con swagger
// @version 1.0
// @description Backend de prueba enfocado en guiar el desarrollo
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

func main() {

	time.Local = time.UTC
	// Cargar Variables de entorno:
	util.LoadEnv()

	// Log
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Start template-go-rest")
	log.Printf("serverUp, %s", os.Getenv("ADDR"))

	// Cargar base de datos
	model.LoadDatabase()

	//Raiz
	app := gin.Default()
	// CORS
	app.Use(middleware.CorsMiddleware())
	// Url Base
	base := app.Group("/api/v1/")
	base.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	controller.Routes(base)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Servicio no encontrado."})
	})

	http.ListenAndServe(os.Getenv("ADDR"), app)

}
