package entities

import (
	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name  string    `gorm:"type:varchar(100);not null" json:"name"`
	Price float64   `gorm:"type:numeric;not null" json:"price"`

	TenantID uuid.UUID `gorm:"type:uuid;null;index" json:"tenant_id"`
	Tenant   Tenant    `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tenant"`

	Timestamp
}
