package repository

import (
	"context"
	"errors"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error)
		GetById(ctx context.Context, tx *gorm.DB, productId string) (entities.Product, error)
		GetByName(ctx context.Context, tx *gorm.DB, name string) (entities.Product, error)
		CheckName(ctx context.Context, tx *gorm.DB, name string) (entities.Product, bool, error)
		Update(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error)
		Delete(ctx context.Context, tx *gorm.DB, productId string) error
	}

	productRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetById(ctx context.Context, tx *gorm.DB, productId string) (entities.Product, error) {
	if tx == nil {
		tx = r.db
	}

	var product entities.Product
	if err := tx.WithContext(ctx).Where("id = ?", productId).Take(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetByName(ctx context.Context, tx *gorm.DB, name string) (entities.Product, error) {
	if tx == nil {
		tx = r.db
	}

	var product entities.Product
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *productRepository) Update(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *productRepository) Delete(ctx context.Context, tx *gorm.DB, productId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entities.Product{}, "id = ?", productId).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) CheckName(ctx context.Context, tx *gorm.DB, name string) (entities.Product, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var product entities.Product
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Product{}, false, nil
		}
		return entities.Product{}, false, err
	}

	return product, true, nil
}
