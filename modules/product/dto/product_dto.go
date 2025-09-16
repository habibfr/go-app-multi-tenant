package dto

import (
	"errors"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/dto"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_PRODUCT   = "failed create product"
	MESSAGE_FAILED_GET_LIST_PRODUCT   = "failed get list product"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_GET_PRODUCT        = "failed get product"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_UPDATE_PRODUCT     = "failed update product"
	MESSAGE_FAILED_DELETE_PRODUCT     = "failed delete product"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL       = "failed verify email"

	// Success
	MESSAGE_SUCCESS_REGISTER_PRODUCT        = "success create product"
	MESSAGE_SUCCESS_GET_LIST_PRODUCT        = "success get list product"
	MESSAGE_SUCCESS_GET_PRODUCT             = "success get product"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_PRODUCT          = "success update product"
	MESSAGE_SUCCESS_DELETE_PRODUCT          = "success delete product"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
)

var (
	ErrCreateProduct          = errors.New("failed to create product")
	ErrGetProductById         = errors.New("failed to get product by id")
	ErrGetProductByName       = errors.New("failed to get product by name")
	ErrProductAlreadyExists   = errors.New("product already exist")
	ErrUpdateProduct          = errors.New("failed to update product")
	ErrProductNotFound        = errors.New("product not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteProduct          = errors.New("failed to delete product")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrNameAlreadyExists      = errors.New("name already exists")
)

type (
	ProductCreateRequest struct {
		Name  string  `json:"name" form:"name" binding:"required,min=2,max=100"`
		Price float64 `json:"price" form:"price" binding:"required,gt=0"`
		// TenantID akan diisi dari context (middleware)
		TenantID string `json:"tenant_id" form:"tenant_id" binding:"-"`
	}

	ProductResponse struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
	}

	ProductPaginationResponse struct {
		Data []ProductResponse `json:"data"`
		dto.PaginationResponse
	}

	GetAllProductRepositoryResponse struct {
		Products []entities.Product `json:"products"`
		dto.PaginationResponse
	}

	ProductUpdateRequest struct {
		Name  string  `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		Price float64 `json:"price" form:"price" binding:"omitempty,gt=0"`
	}

	ProductUpdateResponse struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
	}
)
