package service

import (
	"tpm_6_HendriHeryanto/models"
	"tpm_6_HendriHeryanto/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepo
}

func (p *ProductService) Get() ([]*models.Product, error) {
	return p.ProductRepo.GetProduct()
}

func (p *ProductService) Create(product *models.Product) (*models.Product, error) {
	return p.ProductRepo.CreateProduct(product)
}

func (p *ProductService) Update(id uint64, product *models.Product) (*models.Product, error) {
	return p.ProductRepo.UpdateProduct(id, product)
}

func (p *ProductService) Delete(id uint64) error {
	return p.ProductRepo.DeleteProduct(id)
}
