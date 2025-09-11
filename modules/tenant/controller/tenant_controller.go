package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/query"
	"github.com/Caknoooo/go-gin-clean-starter/modules/tenant/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/Caknoooo/go-pagination"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	TenantController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	tenantController struct {
		tenantService service.TenantService
		db            *gorm.DB
	}
)

func NewTenantController(injector *do.Injector, ts service.TenantService) TenantController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	return &tenantController{
		tenantService: ts,
		db:            db,
	}
}

func (c *tenantController) Create(ctx *gin.Context) {
	var tenant dto.TenantCreateRequest
	if err := ctx.ShouldBind(&tenant); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.tenantService.Create(ctx.Request.Context(), tenant)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_TENANT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_TENANT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *tenantController) GetById(ctx *gin.Context) {
	tenantId := ctx.Param("id")

	result, err := c.tenantService.GetTenantById(ctx.Request.Context(), tenantId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TENANT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TENANT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *tenantController) GetAll(ctx *gin.Context) {
	var filter = &query.TenantFilter{}
	filter.BindPagination(ctx)

	ctx.ShouldBindQuery(filter)

	tenants, total, err := pagination.PaginatedQueryWithIncludable[query.Tenant](c.db, filter)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TENANT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	paginationResponse := pagination.CalculatePagination(filter.Pagination, total)
	response := pagination.NewPaginatedResponse(http.StatusOK, dto.MESSAGE_SUCCESS_GET_LIST_TENANT, tenants, paginationResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *tenantController) Update(ctx *gin.Context) {
	tenantId := ctx.Param("id")
	var tenant dto.TenantUpdateRequest
	if err := ctx.ShouldBind(&tenant); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.tenantService.Update(ctx.Request.Context(), tenant, tenantId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TENANT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TENANT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *tenantController) Delete(ctx *gin.Context) {
	tenantId := ctx.Param("id")

	if err := c.tenantService.Delete(ctx.Request.Context(), tenantId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TENANT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TENANT, nil)
	ctx.JSON(http.StatusOK, res)
}
