package entities

import (
	"github.com/google/uuid"
)

type Tenant struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(100);not null" json:"name"`

	Products []Product `gorm:"foreignKey:TenantID" json:"products"`
	Users    []User    `gorm:"foreignKey:TenantID" json:"users"`
	Timestamp
}
