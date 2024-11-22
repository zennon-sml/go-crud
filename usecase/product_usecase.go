package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-crud/model"
	"github.com/zennon-sml/go-crud/repository"
)

type ProductUseCase struct {
	repository repository.ProdcutRepository
}

func NewProductUseCase(repo repository.ProdcutRepository) ProductUseCase {
	return ProductUseCase{
    repository: repo,
  }
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error){
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) DeleteProductById(id int) (productId int, err error) {
	//call the repository
	deletedId, err := pu.DeleteProductById(id) 
	if err != nil {
		return 0, err
	}

	return deletedId, nil
}