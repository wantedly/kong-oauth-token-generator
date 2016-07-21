package server

import (
	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/controller"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func Run(config *config.KongConfiguration) {
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())
	client := kong.NewClient(nil, config)
	apiController := controller.NewAPIController(client)
	oauth2Controller := controller.NewOAuth2Controller(client)

	r.GET("/", apiController.Index)

	r.GET("/apis", apiController.Index)
	r.POST("/apis", apiController.Create)
	r.GET("/apis/:apiName", apiController.Get)
	r.POST("/apis/:apiName/delete", apiController.Delete)
	r.GET("/new-api", apiController.New)

	r.GET("/oauth2s", oauth2Controller.Index)
	r.POST("/oauth2s", oauth2Controller.Create)
	r.GET("/oauth2s/:consumerName", oauth2Controller.Get)
	r.POST("/oauth2s/:consumerName/delete", oauth2Controller.Delete)
	r.GET("/new-oauth2", apiController.New)

	r.Run()
}
