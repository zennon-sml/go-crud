package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-crud/model"
	"github.com/zennon-sml/go-crud/usecase"
)

type productController struct {
  productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
  return productController{
    productUseCase: usecase,
  }
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
