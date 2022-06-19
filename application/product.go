package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func (product *Product) IsValid() (bool, error) {
	if product.ID == "" {
		return false, errors.New("id is required")
	}
	if product.Name == "" {
		return false, errors.New("name is required")
	}
	if product.Price == 0 {
		return false, errors.New("price is required")
	}
	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New("status is required")
	}
	return true, nil
}

func (product *Product) Enable() error {
	if product.Status == ENABLED {
		return errors.New("Product is already enabled")
	}
	product.Status = ENABLED
	return nil
}

func (product *Product) Disable() error {
	if product.Status == DISABLED {
		return errors.New("Product is already disabled")
	}
	product.Status = DISABLED
	return nil
}

func (product *Product) GetId() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetStatus() string {
	return product.Status
}

func (product *Product) GetPrice() float64 {
	return product.Price
}
