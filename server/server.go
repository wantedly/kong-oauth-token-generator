package server

import (
	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/controller"
)

func Run(config *config.KongConfiguration) {
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("templates/*")

	rootController := controller.NewRootController(config)

	r.GET("/", rootController.Index)

	r.Run()
}
