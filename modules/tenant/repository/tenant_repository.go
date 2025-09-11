package repository

import (
	"context"
	"errors"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

type (
	TenantRepository interface {
		Create(ctx context.Context, tx *gorm.DB, tenant entities.Tenant) (entities.Tenant, error)
		GetById(ctx context.Context, tx *gorm.DB, tenantId string) (entities.Tenant, error)
		GetByName(ctx context.Context, tx *gorm.DB, name string) (entities.Tenant, error)
		// CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entities.Tenant, bool, error)
		CheckName(ctx context.Context, tx *gorm.DB, name string) (entities.Tenant, bool, error)
		Update(ctx context.Context, tx *gorm.DB, tenant entities.Tenant) (entities.Tenant, error)
		Delete(ctx context.Context, tx *gorm.DB, tenantId string) error
	}

	tenantRepository struct {
		db *gorm.DB
	}
)

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		db: db,
	}
}

func (r *tenantRepository) Create(ctx context.Context, tx *gorm.DB, tenant entities.Tenant) (entities.Tenant, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&tenant).Error; err != nil {
		return entities.Tenant{}, err
	}

	return tenant, nil
}

func (r *tenantRepository) GetById(ctx context.Context, tx *gorm.DB, tenantId string) (entities.Tenant, error) {
	if tx == nil {
		tx = r.db
	}

	var tenant entities.Tenant
	if err := tx.WithContext(ctx).Where("id = ?", tenantId).Take(&tenant).Error; err != nil {
		return entities.Tenant{}, err
	}

	return tenant, nil
}

func (r *tenantRepository) GetByName(ctx context.Context, tx *gorm.DB, name string) (entities.Tenant, error) {
	if tx == nil {
		tx = r.db
	}

	var tenant entities.Tenant
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&tenant).Error; err != nil {
		return entities.Tenant{}, err
	}

	return tenant, nil
}

func (r *tenantRepository) Update(ctx context.Context, tx *gorm.DB, tenant entities.Tenant) (entities.Tenant, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&tenant).Error; err != nil {
		return entities.Tenant{}, err
	}

	return tenant, nil
}

func (r *tenantRepository) Delete(ctx context.Context, tx *gorm.DB, tenantId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entities.Tenant{}, "id = ?", tenantId).Error; err != nil {
		return err
	}

	return nil
}

func (r *tenantRepository) CheckName(ctx context.Context, tx *gorm.DB, name string) (entities.Tenant, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var tenant entities.Tenant
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Tenant{}, false, nil
		}
		return entities.Tenant{}, false, err
	}

	return tenant, true, nil
}
