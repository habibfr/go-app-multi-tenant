package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/product/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product/query"
	"github.com/Caknoooo/go-gin-clean-starter/modules/product/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/Caknoooo/go-pagination"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	ProductController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	productController struct {
		productService service.ProductService
		db             *gorm.DB
	}
)

func NewProductController(injector *do.Injector, ps service.ProductService) ProductController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	return &productController{
		productService: ps,
		db:             db,
	}
}

func (c *productController) Create(ctx *gin.Context) {
	var product dto.ProductCreateRequest
	if err := ctx.ShouldBind(&product); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	tenantID := ctx.MustGet("tenant_id").(string)
	product.TenantID = tenantID // set tenantID dari middleware ke request

	result, err := c.productService.Create(ctx.Request.Context(), product)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_PRODUCT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_PRODUCT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetById(ctx *gin.Context) {
	productId := ctx.Param("id")

	result, err := c.productService.GetProductById(ctx.Request.Context(), productId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PRODUCT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetAll(ctx *gin.Context) {
	tenantId := ctx.MustGet("tenant_id").(string)
	var filter = &query.ProductFilter{
		TenantID: tenantId,
		Name:     ctx.Query("name"),
	}
	filter.BindPagination(ctx)

	ctx.ShouldBindQuery(filter)

	products, total, err := pagination.PaginatedQueryWithIncludable[query.Product](c.db, filter)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	paginationResponse := pagination.CalculatePagination(filter.Pagination, total)
	response := pagination.NewPaginatedResponse(http.StatusOK, dto.MESSAGE_SUCCESS_GET_LIST_PRODUCT, products, paginationResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) Update(ctx *gin.Context) {
	productId := ctx.Param("id")
	var product dto.ProductUpdateRequest
	if err := ctx.ShouldBind(&product); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.productService.Update(ctx.Request.Context(), product, productId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PRODUCT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PRODUCT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) Delete(ctx *gin.Context) {
	productId := ctx.Param("id")

	if err := c.productService.Delete(ctx.Request.Context(), productId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_PRODUCT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_PRODUCT, nil)
	ctx.JSON(http.StatusOK, res)
}
