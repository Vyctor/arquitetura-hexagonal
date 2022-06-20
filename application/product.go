package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

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
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,required"`
	Status string  `valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (product *Product) IsValid() (bool, error) {
	if product.Status == "" {
		product.Status = DISABLED
	}
	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New("status is required")
	}
	if product.ID == "" {
		return false, errors.New("id is required")
	}
	if product.Name == "" {
		return false, errors.New("name is required")
	}
	if product.Price == 0 {
		return false, errors.New("price is required")
	}
	if product.Price <= 0 {
		return false, errors.New("price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(product)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (product *Product) Enable() error {
	if product.Price <= 0 {
		return errors.New("price must be greater than zero to enable the product")
	}

	if product.Status == ENABLED {
		return errors.New("product is already enabled")
	}

	product.Status = ENABLED
	return nil
}

func (product *Product) Disable() error {
	if product.Price != 0 {
		return errors.New("price must be zero in order to have the product disabled")
	}

	if product.Status == DISABLED {
		return errors.New("product is already disabled")
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
