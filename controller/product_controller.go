package controller

import (
	"net/http"
	"strconv"

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

  products, err := p.productUseCase.GetProducts()
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, err)
  }

  ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

  var product model.Product
  err := ctx.BindJSON(&product)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }

  insertedProduct, err := p.productUseCase.CreateProduct(product)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, err)
    return
  }

  ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

  id := ctx.Param("id")
  // check if param isnt null
  if id == "" {
    response := model.Response {
      Message: "Id cannot be null",
    }
    ctx.JSON(http.StatusBadRequest, response)
    return
  }

  productId, err := strconv.Atoi(id) 
  //check conversion error
  if err != nil{
    response := model.Response {
      Message: "conversion from string to int went wrong",
    }
    ctx.JSON(http.StatusBadRequest, response)
    return
  }

  product, err := p.productUseCase.GetProductById(productId)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, err)
  }

  if product == nil {
    response := model.Response {
      Message: "Product not found on db",
    }
    ctx.JSON(http.StatusNotFound, response)
    return
  }

  ctx.JSON(http.StatusOK, product)
}