package query

import (
	"github.com/Caknoooo/go-pagination"
	"gorm.io/gorm"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductFilter struct {
	pagination.BaseFilter
	Name     string `form:"name"`      // tambahkan ini
	TenantID string `form:"tenant_id"` // tambahkan ini
}

func (f *ProductFilter) ApplyFilters(query *gorm.DB) *gorm.DB {
	// Apply your filters here
	if f.Name != "" {
		query = query.Where("name ILIKE ?", "%"+f.Name+"%")
	}
	if f.TenantID != "" {
		query = query.Where("tenant_id = ?", f.TenantID)
	}
	return query
}

func (f *ProductFilter) GetTableName() string {
	return "products"
}

func (f *ProductFilter) GetSearchFields() []string {
	return []string{"name"}
}

func (f *ProductFilter) GetDefaultSort() string {
	return "id asc"
}

func (f *ProductFilter) GetIncludes() []string {
	return f.Includes
}

func (f *ProductFilter) GetPagination() pagination.PaginationRequest {
	return f.Pagination
}

func (f *ProductFilter) Validate() {
	var validIncludes []string
	allowedIncludes := f.GetAllowedIncludes()
	for _, include := range f.Includes {
		if allowedIncludes[include] {
			validIncludes = append(validIncludes, include)
		}
	}
	f.Includes = validIncludes
}

func (f *ProductFilter) GetAllowedIncludes() map[string]bool {
	return map[string]bool{}
}
