package main

import (
	"interview/internal/controllers"
	"interview/internal/db/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.New()
	db.MigrateDatabase()

	ginEngine := gin.Default()

	var taxController controllers.TaxController
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
