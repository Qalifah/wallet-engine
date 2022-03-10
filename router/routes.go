package router

import (
	"github.com/gin-gonic/gin"
	"wallet-engine/handler"
)

func Init(ctrl *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	wallet := router.Group("/wallet")
	wallet.POST("/new", ctrl.CreateNewWallet)
	wallet.POST("/debit", ctrl.DebitWallet)
	wallet.POST("/credit", ctrl.CreditWallet)
	wallet.POST("/activate", ctrl.ActivateWallet)
	wallet.POST("/deactivate", ctrl.DeActivateWallet)

	return router
}
