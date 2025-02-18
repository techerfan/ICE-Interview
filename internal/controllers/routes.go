package controllers

import "github.com/gin-gonic/gin"

func (t *CartController) setupRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/", t.ShowAddItemForm)
	ginEngine.POST("/add-item", t.AddItem)
	ginEngine.GET("/remove-cart-item", t.DeleteCartItem)
}
