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
// @Router       /comingTableProduct [post]
// @Summary      create comingTableProduct
// @Description  api for create comingTableProduct
// @Tags         comingTableProducts
// @Accept       json
// @Produce      json
// @Param        comingTableProducts    body     models.CreateComingTableProduct  true  "date of comingTableProduct"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateComingTableProduct(c *gin.Context) {
	fmt.Println("Method POST")
	var comingTableProduct models.CreateComingTableProduct
	err := c.ShouldBind(&comingTableProduct)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.ComingTableProduct().CreateComingTableProduct(comingTableProduct)
	if err != nil {
		h.log.Error("error CreateComingTableProduct", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateComingTableProduct")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /comingTableProduct/{id} [get]
// @Summary      get comingTableProduct
// @Description  get comingTableProduct
// @Tags         comingTableProducts
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTableProduct"  Format(uuid)
// @Success      200  {object}   models.ComingTableProduct
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetComingTableProduct(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.ComingTableProduct().GetComingTableProduct(models.IdRequestComingTableProduct{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error ComingTableProduct Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTableProduct/{id} [put]
// @Summary      updateda comingTableProduct
// @Description   api fot update comingTableProduct
// @Tags         comingTableProducts
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTableProduct"
// @Param        staff    body     models.CreateComingTableProduct  true  "id of comingTableProduct"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateComingTableProduct(c *gin.Context) {
	fmt.Println("Method PUT")
	var comingTableProduct models.ComingTableProduct
	err := c.ShouldBind(&comingTableProduct)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	comingTableProduct.Id = c.Param("id")
	resp, err := h.strg.ComingTableProduct().UpdateComingTableProduct(comingTableProduct)
	if err != nil {
		fmt.Println("error ComingTableProduct Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTableProduct/{id} [delete]
// @Summary      delete comingTableProduct
// @Description   api fot delete comingTableProduct
// @Tags         comingTableProducts
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTableProduct"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteComingTableProduct(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error ComingTableProduct Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.ComingTable().DeleteComingTable(models.IdRequestComingTable{Id: id})
	if err != nil {
		h.log.Error("error ComingTableProduct Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTableProduct [get]
// @Summary      List comingTableProduct
// @Description  get comingTableProduct
// @Tags         comingTableProducts
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllComingTableProduct
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllComingTableProduct(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllComingTablesProduct")
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

	resp, err := h.strg.ComingTableProduct().GetAllComingTableProduct(models.GetAllComingTableProductRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error ComingTableProduct GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllComingTableProduct")
	c.JSON(http.StatusOK, resp)
}