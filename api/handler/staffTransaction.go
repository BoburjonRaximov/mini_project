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
// @Router       /staffTransaction [post]
// @Summary      create staffTransaction
// @Description  api for create staffTransaction
// @Tags         staffTransactions
// @Accept       json
// @Produce      json
// @Param        staffTransaction    body     models.CreateStaffTransaction  true  "date of staffTransaction"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaffTransaction(c *gin.Context) {
	fmt.Println("Method POST")
	var staffTransaction models.CreateStaffTransaction
	err := c.ShouldBindJSON(&staffTransaction)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.StaffTransaction().CreateStaffTransaction(c.Request.Context(), staffTransaction)
	if err != nil {
		h.log.Error("error CreateStaffTransaction", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateStaffTransaction")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /staffTransaction/{id} [get]
// @Summary      get staffTransaction
// @Description  get staffTransaction
// @Tags         staffTransactions
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTransaction"  Format(uuid)
// @Success      200  {object}   models.StaffTransaction
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaffTransaction(c *gin.Context) {
	fmt.Println("Method GET")
	staffTransaction := models.StaffTransaction{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, staffTransaction)
	if err != nil {
		fmt.Println("not found redis in cache")
	}
	if ok {
		c.JSON(http.StatusOK, staffTransaction)
		return
	}

	resp, err := h.strg.StaffTransaction().GetStaffTransaction(c.Request.Context(), models.IdRequestStaffTransaction{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error StaffTransaction Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, 5*time.Minute)
	if err != nil {
		fmt.Println("error create redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staffTransaction/{id} [put]
// @Summary      updateda staffTransaction
// @Description   api fot update staffTransaction
// @Tags         staffTransactions
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTransaction"
// @Param        staff    body     models.CreateStaff  true  "id of staffTransaction"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffTransaction(c *gin.Context) {
	fmt.Println("Method PUT")
	var staffTransaction models.StaffTransaction
	err := c.ShouldBindJSON(&staffTransaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staffTransaction.Id = c.Param("id")
	resp, err := h.strg.StaffTransaction().UpdateStaffTransaction(c.Request.Context(), staffTransaction)
	if err != nil {
		fmt.Println("error StaffTransaction Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Delete(c.Request.Context(), staffTransaction.Id)
	if err != nil {
		fmt.Println("error delete redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staffTransaction/{id} [delete]
// @Summary      delete staffTransaction
// @Description   api fot delete staffTransaction
// @Tags         staffTransactions
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTransaction"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaffTransaction(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error StaffTransaction Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.StaffTransaction().DeleteStaffTransaction(c.Request.Context(), models.IdRequestStaffTransaction{Id: id})
	if err != nil {
		h.log.Error("error StaffTransaction Delete:", logger.Error(err))
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
// @Router       /staffTransaction [get]
// @Summary      List staffTransactions
// @Description  get staffTransaction
// @Tags         staffTransactions
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllStaffTransaction
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaffTransaction(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllStaffTransaction")
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

	resp, err := h.strg.StaffTransaction().GetAllStaffTransaction(c.Request.Context(), models.GetAllStaffTransactionRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error StaffTransaction GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllStaffTransaction")
	c.JSON(http.StatusOK, resp)

}
