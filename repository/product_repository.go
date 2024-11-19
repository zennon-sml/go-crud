package repository

import (
	"database/sql"
	"fmt"

	"github.com/zennon-sml/go-crud/model"
)

type ProdcutRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProdcutRepository {
	return ProdcutRepository{
    connection: connection,
  }
}

func (pr *ProdcutRepository) GetProducts() ([]model.Product, error) {
  query := "SELECT * FROM product"
  rows, err := pr.connection.Query(query)
  if err != nil {
    fmt.Println(err)
    return []model.Product{}, err
  }

  var productList []model.Product
  var productObj model.Product

  for rows.Next() {
    err = rows.Scan(
      &productObj.ID,
      &productObj.Name,
      &productObj.Price)
    if err != nil {
      fmt.Println(err)
      return []model.Product{}, err
    }
  }

  productList = append(productList, productObj)

  return  productList, nil
}