package controller

import (
	"fmt"
	"ltinfo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	cs services.CurrencyService
}

func NewCurrencyController() *CurrencyController {
	return &CurrencyController{
		cs: services.NewCurrencyService(),
	}
}

func (curr *CurrencyController) GetCurrencyDetails(ctx *gin.Context) {
	symbol := ctx.Params.ByName("symbol")
	if symbol == "all" {
		fmt.Println("symbol", symbol == "all", symbol)
		response := curr.cs.GetAllCurrency()
		ctx.JSON(http.StatusOK, gin.H{
			"currencies": response,
		})
	} else {
		fmt.Println("symbol", symbol)
		response := curr.cs.GetCurrencyDetails(symbol)
		ctx.JSON(http.StatusOK, response)
	}

}
