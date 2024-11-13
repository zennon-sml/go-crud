package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/zennon-sml/go-crud/model"
)

type productController struct {

}

func NewProductController() productController {
  return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
  products := []model.Product{
    {
      ID: 1,
      Name: "Batata frita",
      Price: 20,
    },
  }

  ctx.JSON(http.StatusOK, products)
}
