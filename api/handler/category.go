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
// @Router       /category [post]
// @Summary      create category
// @Description  api for create category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category    body     models.CreateCategory  true  "date of category"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateCategory(c *gin.Context) {
	fmt.Println("Method POST")
	var category models.CreateCategory
	err := c.ShouldBind(&category)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.Category().CreateCategory(category)
	if err != nil {
		h.log.Error("error CreateCategory", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateCategory")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /category/{id} [get]
// @Summary      get category
// @Description  get category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of category"  Format(uuid)
// @Success      200  {object}   models.Category
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetCategory(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.Category().GetCategory(models.IdRequestCategory{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Category Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /category/{id} [put]
// @Summary      updateda category
// @Description   api fot update category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of category"
// @Param        sale    body     models.CreateCategory  true  "id of category"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateCategory(c *gin.Context) {
	fmt.Println("Method PUT")
	var category models.Category
	err := c.ShouldBind(&category)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	category.Id = c.Param("id")
	resp, err := h.strg.Category().UpdateCategory(category)
	if err != nil {
		fmt.Println("error Category Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /category/{id} [delete]
// @Summary      delete category
// @Description   api fot delete category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of category"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteCategory(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Category Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Category().DeleteCategory(models.IdRequestCategory{Id: id})
	if err != nil {
		h.log.Error("error Category Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /category [get]
// @Summary      List category
// @Description  get category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllCategory
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllCategory(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllCategory")
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

	resp, err := h.strg.Category().GetAllCategory(models.GetAllCategoryRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Category GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllCategory")
	c.JSON(http.StatusOK, resp)
}
