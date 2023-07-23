package router

import (
	"ltinfo/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	routes := gin.Default()

	// Ping test
	routes.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "test2")
	})

	currencyController := controller.NewCurrencyController()
	CurrencyRoutes(routes, *currencyController)

	return routes
}

func CurrencyRoutes(route *gin.Engine, controller controller.CurrencyController) {
	route.GET("/currency/:symbol", controller.GetCurrencyDetails)

}
