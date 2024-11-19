package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

  products, err := p.productUseCase.GetProducts()
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, err)
  }

  ctx.JSON(http.StatusOK, products)
}
