package route

import "github.com/gin-gonic/gin"

func HandleRequest() {
	rte := gin.Default()
	ReceiptRoute(rte)

	rte.Run(":8080")
}
