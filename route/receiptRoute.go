package route

import (
	"cooking-receipt/controller"
	"cooking-receipt/wrapper/receiptWrapper"
	"github.com/gin-gonic/gin"
)

func ReceiptRoute(r *gin.Engine) {
	routes := r.Group("/api/v1/recipes")

	var receiptwrapper = receiptWrapper.GetInstance()
	var receiptController = controller.ReceiptController{Receipt: receiptwrapper}

	routes.GET("/", receiptController.FindAll)
	routes.GET("/:id", receiptController.FindRecipe)
	routes.POST("/", receiptController.Create)
	routes.PUT("/:id", receiptController.Update)
	routes.DELETE("/:id", receiptController.Delete)
}
