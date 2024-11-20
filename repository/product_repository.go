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
  query := "SELECT id, name, price FROM products"
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
    productList = append(productList, productObj)
  }


  return  productList, nil
}

func (pr *ProdcutRepository) CreateProduct(product model.Product) (int, error) {

  var id int
  query, err := pr.connection.Prepare("INSERT INTO products" + 
  "(name, price) " +
  "VALUES ($1, $2) RETURNING id")
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  err = query.QueryRow(product.Name, product.Price).Scan(&id)
  if err != nil {
    fmt.Println(err)
    return 0, err
  }

  query.Close()
  return id, nil
}

func (pr *ProdcutRepository) GetProductById(id int) (*model.Product, error) {
  query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1")
  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  var product model.Product

  err = query.QueryRow(id).Scan(
    &product.ID,
    &product.Name,
    &product.Price,
  )
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, nil
    }

    return nil, err
  }

  query.Close()
  return &product, nil
}