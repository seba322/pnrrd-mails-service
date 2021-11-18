package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/citiaps/pnrrd-formulario-jerarquias/controller"
	"github.com/citiaps/pnrrd-formulario-jerarquias/middleware"
	"github.com/citiaps/pnrrd-formulario-jerarquias/model"
	"github.com/citiaps/pnrrd-formulario-jerarquias/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/citiaps/pnrrd-formulario-jerarquias/docs"
)

// @title Documentación Servicio formularios Pnrrd
// @version 1.0
// @description Backend enfocado de formulario de jarquías para Pnrrd
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1

func main() {

	time.Local = time.UTC
	// Cargar Variables de entorno:
	util.LoadEnv()

	// Log
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Start pnrrd-formulario-jerarquias")
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
