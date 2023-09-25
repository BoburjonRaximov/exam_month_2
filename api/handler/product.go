package handler

import (
	"errors"
	"fmt"
	"net/http"
	"new_project/models"
	"new_project/pkg/helper"
	"new_project/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Router       /product [post]
// @Summary      create product
// @Description  api for create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product    body     models.CreateProduct  true  "date of product"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateProduct(c *gin.Context) {
	fmt.Println("Method POST")
	var product models.CreateProduct
	err := c.ShouldBind(&product)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.Product().CreateProduct(product)
	if err != nil {
		h.log.Error("error CreateProduct", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateProduct")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /product/{id} [get]
// @Summary      get product
// @Description  get product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of product"  Format(uuid)
// @Success      200  {object}   models.Product
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetProduct(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.Product().GetProduct(models.IdRequestProduct{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Product Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /product/{id} [put]
// @Summary      updateda product
// @Description   api fot update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of product"
// @Param        product    body     models.CreateProduct  true  "id of product"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateProduct(c *gin.Context) {
	fmt.Println("Method PUT")
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	product.Id = c.Param("id")
	resp, err := h.strg.Product().UpdateProduct(product)
	if err != nil {
		fmt.Println("error Product Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /product/{id} [delete]
// @Summary      delete product
// @Description   api fot delete product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of products"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteProduct(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Product Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Product().DeleteProduct(models.IdRequestProduct{Id: id})
	if err != nil {
		h.log.Error("error Product Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /product [get]
// @Summary      List product
// @Description  get product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        limit     query     integer  true   "limit for response"  Default(10)
// @Param        page      query     integer  true   "page of req"  Default(1)
// @Param        search    query     string   false  "search name"
// @Param        barcode   query     string   false  "search barcode"
// @Success      200  {object}  models.GetAllProduct
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllProduct(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllPRODUCT")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.strg.Product().GetAllProduct(models.GetAllProductRequest{
		Page:   page,
		Limit:  limit,
		SearchName: c.Query("search"),
		Barcode: c.Query("barcode"),
	})
	if err != nil {
		h.log.Error("error product GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllProduct")
	c.JSON(http.StatusOK, resp)
}
