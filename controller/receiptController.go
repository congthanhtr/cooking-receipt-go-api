package controller

import (
	"cooking-receipt/model"
	"cooking-receipt/wrapper/receiptWrapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReceiptController struct {
	Receipt receiptWrapper.IReceiptWrapper
}

func (controller *ReceiptController) FindAll(c *gin.Context) {
	recipes, err := controller.Receipt.Find()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

func (controller *ReceiptController) FindRecipe(c *gin.Context) {
	recipeId := c.Params.ByName("id")

	recipe, err := controller.Receipt.FindById(recipeId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (controller *ReceiptController) Create(c *gin.Context) {
	err := controller.Receipt.Create(model.CookingReceipt{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (controller *ReceiptController) Update(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	receipt, err := controller.Receipt.Save(recipeId, model.CookingReceipt{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, receipt)
}

func (controller *ReceiptController) Delete(c *gin.Context) {
	recipeId := c.Params.ByName("id")
	err := controller.Receipt.Delete(recipeId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (controller *ReceiptController) Seach(c *gin.Context) {
	search := c.Params.ByName("search")
	receipt, err := controller.Receipt.Search(search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, receipt)
}
