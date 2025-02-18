package main

import (
	"interview/internal/controllers"
	"interview/internal/db/mysql"
	"interview/internal/db/redis/productredis"
	"interview/internal/service/cart"
	"interview/internal/service/product"
	"interview/internal/validator/cartvalidator"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.New()
	db.MigrateDatabase()

	productRedis := productredis.New(productredis.Config{
		Host: "127.0.0.1",
		Port: 4000,
		DB:   1,
	})

	productService := product.New(productRedis)

	cartService := cart.New(db, productService)

	cartValidator := cartvalidator.New(productRedis, db)

	ginEngine := gin.Default()

	cartController := controllers.New(cartService, cartValidator, ginEngine)
	cartController.Start(":8088")
}
