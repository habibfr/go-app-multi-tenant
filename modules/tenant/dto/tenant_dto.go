package dto

import (
	"errors"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/dto"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_TENANT    = "failed create tenant"
	MESSAGE_FAILED_GET_LIST_TENANT    = "failed get list tenant"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_GET_TENANT         = "failed get tenant"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_UPDATE_TENANT      = "failed update tenant"
	MESSAGE_FAILED_DELETE_TENANT      = "failed delete tenant"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL       = "failed verify email"

	// Success
	MESSAGE_SUCCESS_REGISTER_TENANT         = "success create tenant"
	MESSAGE_SUCCESS_GET_LIST_TENANT         = "success get list tenant"
	MESSAGE_SUCCESS_GET_TENANT              = "success get tenant"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_TENANT           = "success update tenant"
	MESSAGE_SUCCESS_DELETE_TENANT           = "success delete tenant"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
)

var (
	ErrCreateTenant           = errors.New("failed to create tenant")
	ErrGetTenantById          = errors.New("failed to get tenant by id")
	ErrGetTenantByEmail       = errors.New("failed to get tenant by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateTenant           = errors.New("failed to update tenant")
	ErrTenantNotFound         = errors.New("tenant not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteTenant           = errors.New("failed to delete tenant")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrNameAlreadyExists      = errors.New("name already exists")
)

type (
	TenantCreateRequest struct {
		Name string `json:"name" form:"name" binding:"required,min=2,max=100"`
	}

	TenantResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	TenantPaginationResponse struct {
		Data []TenantResponse `json:"data"`
		dto.PaginationResponse
	}

	GetAllTenantRepositoryResponse struct {
		Tenants []entities.Tenant `json:"tenants"`
		dto.PaginationResponse
	}

	TenantUpdateRequest struct {
		Name string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
	}

	TenantUpdateResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
