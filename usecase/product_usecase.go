package usecase

import "github.com/zennon-sml/go-crud/model"

type ProductUseCase struct {
	//repository
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}