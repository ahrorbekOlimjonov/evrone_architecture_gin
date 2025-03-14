package v1

import (
	"ai-seller/internal/entity"
	"ai-seller/internal/usecase"
	"ai-seller/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productRoutes struct {
	t usecase.UseCases
	l logger.Interface
	v *validator.Validate
}

func NewProductRoutes(apiV1Group *gin.RouterGroup, t usecase.UseCases, l logger.Interface) {
	p := &productRoutes{t, l, validator.New(validator.WithRequiredStructEnabled())}

	productGroup := apiV1Group.Group("/product")
	{
		productGroup.POST("/", p.createProduct)
		productGroup.GET("/:id", p.getProduct)
		productGroup.PUT("/", p.updateProduct)
		productGroup.DELETE("/:id", p.deleteProduct)
	}
}


// @Summary     Create product
// @Description Create a new product
// @ID          create-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       request body entity.Product true "Product request"
// @Success     201 {object} entity.Product
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /product [post]
func (r *productRoutes) createProduct(ctx *gin.Context) {
	var product entity.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		r.l.Error(err, "http - v1 - createProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := r.t.CreateProduct(ctx, product)
	if err != nil {
		r.l.Error(err, "http - v1 - createProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database problems"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "product created"})
}

// @Summary     Get product
// @Description Get a product by ID
// @ID          get-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       id path string true "Product ID"
// @Success     200 {object} entity.Product
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /product/{id} [get]
func (r *productRoutes) getProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := r.t.GetProduct(ctx, id)
	if err != nil {
		r.l.Error(err, "http - v1 - getProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database problems"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// @Summary     Update product
// @Description Update a product by ID
// @ID          update-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       id path string true "Product ID"
// @Param       request body entity.Product true "Product request"
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /product [put]
func (r *productRoutes) updateProduct(ctx *gin.Context) {
	var product entity.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		r.l.Error(err, "http - v1 - updateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := r.t.UpdateProduct(ctx, product)
	if err != nil {
		r.l.Error(err, "http - v1 - updateProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database problems"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

// @Summary     Delete product
// @Description Delete a product by ID
// @ID          delete-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       id path string true "Product ID"
// @Success     204 {object} nil
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /product/{id} [delete]
func (r *productRoutes) deleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.t.DeleteProduct(ctx, id)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database problems"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
