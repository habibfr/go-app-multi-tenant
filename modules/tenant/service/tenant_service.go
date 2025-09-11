package service

import (
	"context"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantService interface {
	Create(ctx context.Context, req dto.TenantCreateRequest) (dto.TenantResponse, error)
	GetTenantById(ctx context.Context, tenantId string) (dto.TenantResponse, error)
	Update(ctx context.Context, req dto.TenantUpdateRequest, tenantId string) (dto.TenantUpdateResponse, error)
	Delete(ctx context.Context, tenantId string) error
}

type tenantService struct {
	tenantRepository repository.TenantRepository
	db               *gorm.DB
}

func NewTenantService(
	tenantRepo repository.TenantRepository,
	db *gorm.DB,
) TenantService {
	return &tenantService{
		tenantRepository: tenantRepo,
		db:               db,
	}
}

func (s *tenantService) Create(ctx context.Context, req dto.TenantCreateRequest) (dto.TenantResponse, error) {
	_, exists, err := s.tenantRepository.CheckName(ctx, s.db, req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.TenantResponse{}, err
	}
	if exists {
		return dto.TenantResponse{}, dto.ErrNameAlreadyExists
	}

	tenant := entities.Tenant{
		ID:   uuid.New(),
		Name: req.Name,
	}

	createdTenant, err := s.tenantRepository.Create(ctx, s.db, tenant)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	return dto.TenantResponse{
		ID:        createdTenant.ID.String(),
		Name:      createdTenant.Name,
		CreatedAt: createdTenant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdTenant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *tenantService) GetTenantById(ctx context.Context, tenantId string) (dto.TenantResponse, error) {
	tenant, err := s.tenantRepository.GetById(ctx, s.db, tenantId)
	if err != nil {
		return dto.TenantResponse{}, err
	}

	return dto.TenantResponse{
		ID:        tenant.ID.String(),
		Name:      tenant.Name,
		CreatedAt: tenant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tenant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *tenantService) Update(ctx context.Context, req dto.TenantUpdateRequest, tenantId string) (dto.TenantUpdateResponse, error) {
	tenant, err := s.tenantRepository.GetById(ctx, s.db, tenantId)
	if err != nil {
		return dto.TenantUpdateResponse{}, dto.ErrTenantNotFound
	}

	if req.Name != "" {
		tenant.Name = req.Name
	}

	updatedTenant, err := s.tenantRepository.Update(ctx, s.db, tenant)
	if err != nil {
		return dto.TenantUpdateResponse{}, err
	}

	return dto.TenantUpdateResponse{
		ID:        updatedTenant.ID.String(),
		Name:      updatedTenant.Name,
		CreatedAt: updatedTenant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedTenant.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *tenantService) Delete(ctx context.Context, tenantId string) error {
	return s.tenantRepository.Delete(ctx, s.db, tenantId)
}
