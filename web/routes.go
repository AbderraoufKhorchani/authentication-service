package web

import (
	_ "github.com/AbderraoufKhorchani/authentification-service/docs"
	"github.com/AbderraoufKhorchani/authentification-service/internal/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST"}
	config.AllowHeaders = []string{"accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Link"}
	config.AllowCredentials = true
	config.MaxAge = 300
	r.Use(cors.New(config))

	//add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("signup", helpers.Register)
	r.POST("login", helpers.Login)
	r.POST("reset_password", helpers.ResetPassword)
	return r
}
