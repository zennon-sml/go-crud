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
//Quando você usa func (pr *ProductRepository) como o receiver (receptor), está definindo que o método GetProducts pertence à instância do tipo *ProductRepository. Isso permite acessar os campos e métodos da estrutura ProductRepository diretamente dentro de GetProducts.
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

func (pr *ProdcutRepository) DeleteProductById(id int) (productId int, err error) {
  query, err := pr.connection.Prepare("DELETE FROM PRODUCTS WHERE id = $1")
  if err != nil {
    return 0, fmt.Errorf("error on prepare")
  }
  var deletedId int

  err = query.QueryRow(id).Scan(&deletedId)
  if err != nil {
    if err == sql.ErrNoRows {
      return 0, fmt.Errorf("id not found on db")
    }
    return 0, fmt.Errorf("error on the query")
  }

  query.Close()
  return deletedId, nil
}