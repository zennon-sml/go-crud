package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-crud/controller"
	"github.com/zennon-sml/go-crud/usecase"
)

func main() {
  server := gin.Default()

  ProductUseCase := usecase.NewProductUseCase()

  ProductController := controller.NewProductController(ProductUseCase)

  server.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "pong pong",
    })
  })

  server.GET("/products", ProductController.GetProducts)

  server.Run(":8000")
}
