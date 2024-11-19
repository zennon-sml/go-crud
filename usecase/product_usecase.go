package usecase

import (
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