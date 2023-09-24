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
// @Router       /remaining [post]
// @Summary      create remaining
// @Description  api for create remaining
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        remaining    body     models.CreateRemaining  true  "date of remaining"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateRemaining(c *gin.Context) {
	fmt.Println("Method POST")
	var remaining models.CreateRemaining
	err := c.ShouldBindJSON(&remaining)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.Remaining().CreateRemaining(remaining)
	if err != nil {
		h.log.Error("error CreateRemaining", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateRemaining")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /remaining/{id} [get]
// @Summary      get remaining
// @Description  get remaining
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of remaining"  Format(uuid)
// @Success      200  {object}   models.Remaining
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetRemaining(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.Remaining().GetRemaining(models.IdRequestRemaining{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error remaining Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /remaining/{id} [put]
// @Summary      updateda remaining
// @Description   api fot update remaining
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of remaining"
// @Param        staff    body     models.CreateRemaining  true  "id of remaining"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateRemaining(c *gin.Context) {
	fmt.Println("Method PUT")
	var remaining models.Remaining
	err := c.ShouldBindJSON(&remaining)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	remaining.Id = c.Param("id")
	resp, err := h.strg.Remaining().UpdateRemaining(remaining)
	if err != nil {
		fmt.Println("error Remaining Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /remaining/{id} [delete]
// @Summary      delete remaining
// @Description   api fot delete remaining
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of remaining"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteRemaining(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Remaining Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Remaining().DeleteRemaining(models.IdRequestRemaining{Id: id})
	if err != nil {
		h.log.Error("error remaining Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /remaining [get]
// @Summary      List remaining
// @Description  get remaining
// @Tags         remainings
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllRemaining
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllRemaining(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllRemaining")
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

	resp, err := h.strg.Remaining().GetAllRemaining(models.GetAllRemainingRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Remainig GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllRemaining")
	c.JSON(http.StatusOK, resp)

}
