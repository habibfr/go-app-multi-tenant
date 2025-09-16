package service

import (
	"context"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error)
	GetProductById(ctx context.Context, productId string) (dto.ProductResponse, error)
	Update(ctx context.Context, req dto.ProductUpdateRequest, productId string) (dto.ProductUpdateResponse, error)
	Delete(ctx context.Context, productId string) error
}

type productService struct {
	productRepository repository.ProductRepository
	db                *gorm.DB
}

func NewProductService(
	productRepo repository.ProductRepository,
	db *gorm.DB,
) ProductService {
	return &productService{
		productRepository: productRepo,
		db:                db,
	}
}

func (s *productService) Create(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error) {
	_, exists, err := s.productRepository.CheckName(ctx, s.db, req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.ProductResponse{}, err
	}
	if exists {
		return dto.ProductResponse{}, dto.ErrNameAlreadyExists
	}

	product := entities.Product{
		ID:       uuid.New(),
		Name:     req.Name,
		Price:    req.Price,
		TenantID: uuid.MustParse(req.TenantID),
	}

	createdProduct, err := s.productRepository.Create(ctx, s.db, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:        createdProduct.ID.String(),
		Name:      createdProduct.Name,
		Price:     createdProduct.Price,
		CreatedAt: createdProduct.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdProduct.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *productService) GetProductById(ctx context.Context, productId string) (dto.ProductResponse, error) {
	product, err := s.productRepository.GetById(ctx, s.db, productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:        product.ID.String(),
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
		UpdatedAt: product.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *productService) Update(ctx context.Context, req dto.ProductUpdateRequest, productId string) (dto.ProductUpdateResponse, error) {
	product, err := s.productRepository.GetById(ctx, s.db, productId)
	if err != nil {
		return dto.ProductUpdateResponse{}, dto.ErrProductNotFound
	}

	if req.Name != "" {
		product.Name = req.Name
	}

	if req.Price >= 0 {
		product.Price = req.Price
	}

	updatedProduct, err := s.productRepository.Update(ctx, s.db, product)
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	return dto.ProductUpdateResponse{
		ID:        updatedProduct.ID.String(),
		Name:      updatedProduct.Name,
		Price:     updatedProduct.Price,
		CreatedAt: updatedProduct.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedProduct.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *productService) Delete(ctx context.Context, productId string) error {
	return s.productRepository.Delete(ctx, s.db, productId)
}
