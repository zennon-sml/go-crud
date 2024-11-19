package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-crud/controller"
	"github.com/zennon-sml/go-crud/db"
	"github.com/zennon-sml/go-crud/repository"
	"github.com/zennon-sml/go-crud/usecase"
)

func main() {
  server := gin.Default()

  dbConn, err := db.ConnectDB()
  if err != nil {
    panic(err)
  }

  ProdcutRepository := repository.NewProductRepository(dbConn)

  ProductUseCase := usecase.NewProductUseCase(ProdcutRepository)

  ProductController := controller.NewProductController(ProductUseCase)

  server.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "pong pong",
    })
  })

  server.GET("/products", ProductController.GetProducts)

  server.Run(":8000")
}
