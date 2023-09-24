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
// @Router       /comingTable [post]
// @Summary      create comingTable
// @Description  api for create comingTable
// @Tags         comingTables
// @Accept       json
// @Produce      json
// @Param        comingTable    body     models.CreateComingTable  true  "date of comingTable"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateComingTable(c *gin.Context) {
	fmt.Println("Method POST")
	var comingTable models.CreateComingTable
	err := c.ShouldBind(&comingTable)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.ComingTable().CreateComingTable(comingTable)
	if err != nil {
		h.log.Error("error CreateComingTable", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateComingTable")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /comingTable/{id} [get]
// @Summary      get comingTable
// @Description  get comingTable
// @Tags         comingTables
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTables"  Format(uuid)
// @Success      200  {object}   models.ComingTable
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetComingTable(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.ComingTable().GetComingTable(models.IdRequestComingTable{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error ComingTable Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTable/{id} [put]
// @Summary      updateda comingTable
// @Description   api fot update comingTable
// @Tags         comingTables
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTable"
// @Param        staff    body     models.CreateComingTable  true  "id of comingTable"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateComingTable(c *gin.Context) {
	fmt.Println("Method PUT")
	var comingTable models.ComingTable
	err := c.ShouldBind(&comingTable)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	comingTable.Id = c.Param("id")
	resp, err := h.strg.ComingTable().UpdateComingTable(comingTable)
	if err != nil {
		fmt.Println("error ComingTable Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTable/{id} [delete]
// @Summary      delete comingTable
// @Description   api fot delete comingTable
// @Tags         comingTables
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTable"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteComingTable(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error ComingTable Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.ComingTable().DeleteComingTable(models.IdRequestComingTable{Id: id})
	if err != nil {
		h.log.Error("error ComingTable Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /comingTable [get]
// @Summary      List comingTable
// @Description  get comingTable
// @Tags         comingTables
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllComingTable
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllComingTable(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllComingTables")
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

	resp, err := h.strg.ComingTable().GetAllComingTable(models.GetAllComingTableRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error ComingTable GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllComingTable")
	c.JSON(http.StatusOK, resp)
}