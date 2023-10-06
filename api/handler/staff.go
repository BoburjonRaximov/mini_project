package handler

import (
	"errors"
	"fmt"
	"net/http"
	"new_project/models"
	"new_project/pkg/helper"
	"new_project/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Router       /staff [post]
// @Summary      create staff
// @Description  api for create staff
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        staff    body     models.CreateStaff  true  "date of staff"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaff(c *gin.Context) {
	fmt.Println("Method POST")
	var staff models.CreateStaff
	err := c.ShouldBind(&staff)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.Staff().CreateStaff(c.Request.Context(), staff)
	if err != nil {
		h.log.Error("error CreateStaff", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateStaff")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /staff/{id} [get]
// @Summary      get staff
// @Description  get staff
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff"  Format(uuid)
// @Success      200  {object}   models.Staff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaff(c *gin.Context) {
	fmt.Println("Method GET")
	staff := models.Staff{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, staff)
	if err != nil {
		fmt.Println("not found redis in cache")
	}
	if ok {
		c.JSON(http.StatusOK, staff)
		return
	}

	resp, err := h.strg.Staff().GetStaff(c.Request.Context(), models.IdRequestStaff{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Staff Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, 5*time.Minute)
	if err != nil {
		fmt.Println("error create redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staff/{id} [put]
// @Summary      updateda staff
// @Description   api fot update staffs
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff"
// @Param        staff    body     models.CreateStaff  true  "id of staff"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaff(c *gin.Context) {
	fmt.Println("Method PUT")
	var staff models.Staff
	err := c.ShouldBind(&staff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staff.Id = c.Param("id")
	resp, err := h.strg.Staff().UpdateStaff(c.Request.Context(), staff)
	if err != nil {
		fmt.Println("error Staff Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Delete(c.Request.Context(), staff.Id)
	if err != nil {
		fmt.Println("error delete redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staff/{id} [delete]
// @Summary      delete staff
// @Description   api fot delete staff
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaff(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Staff Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Staff().DeleteStaff(c.Request.Context(), models.IdRequestStaff{Id: id})
	if err != nil {
		h.log.Error("error Staff Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error delete redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staff [get]
// @Summary      List staffs
// @Description  get staffs
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllStaff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaff(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllStaffs")
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

	resp, err := h.strg.Staff().GetAllStaff(c.Request.Context(), models.GetAllStaffRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Staff GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllStaffs")
	c.JSON(http.StatusOK, resp)

}
