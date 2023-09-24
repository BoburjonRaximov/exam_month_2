package handler

import (
	"errors"
	"fmt"
	"net/http"
	"new_project/models"
	"new_project/pkg/helper"
	"new_project/pkg/logger"
	"new_project/response"

	"strconv"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Router       /branch [post]
// @Summary      create brach
// @Description  api for create branch
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        branch    body     models.CreateBranch  true  "date of branch"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateBranch(c *gin.Context) {
	var branch models.CreateBranch
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	fmt.Println(h.strg)
	resp, err := h.strg.Branch().CreateBranch(branch)
	if err != nil {
		fmt.Println("error Branch Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})

}

// ListAccounts godoc
// @Router       /branch/{id} [get]
// @Summary      get branch
// @Description  get branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch"  Format(uuid)
// @Success      200  {object}   models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetBranch(c *gin.Context) {
	fmt.Println("MethodGet")

	id := c.Param("id")

	resp, err := h.strg.Branch().GetBranch(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Branch Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /branch/{id} [put]
// @Summary      updateda branch
// @Description   api fot update branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch"
// @Param        branch    body     models.CreateBranch  true  "id of branch"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.Branch
	err := c.ShouldBind(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	branch.Id = c.Param("id")
	resp, err := h.strg.Branch().UpdateBranch(branch)
	if err != nil {
		fmt.Println("error Branch Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

}

// ListAccounts godoc
// @Router       /branch/{id} [delete]
// @Summary      delete branch
// @Description   api fot delete branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Branch Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error Branch Detete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /branch [get]
// @Summary      List branches
// @Description  get branches
// @Tags         branches
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}   models.GetAllBranch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllBranch(c *gin.Context) {
	h.log.Info("request GetAllBranches")
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

	resp, errs := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Page:    page,
		Limit:   limit,
		Search:  c.Query("search"),
		Address: c.Query("address"),
	})
	if errs != nil {
		h.log.Error("error Branch GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllBranches")
	c.JSON(http.StatusOK, resp)
}
