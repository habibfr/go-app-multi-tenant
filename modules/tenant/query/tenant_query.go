package query

import (
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TenantFilter struct {
	pagination.BaseFilter
	Name     string `form:"name"`      // tambahkan ini
	TenantID string `form:"tenant_id"` // tambahkan ini
}

func (f *TenantFilter) ApplyFilters(query *gorm.DB) *gorm.DB {
	// Apply your filters here
	if f.Name != "" {
		query = query.Where("name ILIKE ?", "%"+f.Name+"%")
	}
	if f.TenantID != "" {
		query = query.Where("id = ?", f.TenantID)
	}
	return query
}

func (f *TenantFilter) GetTableName() string {
	return "tenants"
}

func (f *TenantFilter) GetSearchFields() []string {
	return []string{"name"}
}

func (f *TenantFilter) GetDefaultSort() string {
	return "id asc"
}

func (f *TenantFilter) GetIncludes() []string {
	return f.Includes
}

func (f *TenantFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *TenantFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *TenantFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{}
}
