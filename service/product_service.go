package service

import (
	"github.com/marcospmail/devfullcycle-imersao17-goapi/internal/database"
	"github.com/marcospmail/devfullcycle-imersao17-goapi/internal/entity"
)

type ProductService struct {
	ProductDB *database.ProductDB
}

func NewProductService(productDB *database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (productService *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := productService.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (productService *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := productService.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productService *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := productService.ProductDB.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (productService *ProductService) CreateProduct(name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := productService.ProductDB.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
