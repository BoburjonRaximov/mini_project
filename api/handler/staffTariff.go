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
// @Router       /staffTariff [post]
// @Summary      create staffTariff
// @Description  api for create staffTariff
// @Tags         staffTariffs
// @Accept       json
// @Produce      json
// @Param        staffTariff    body     models.CreateStaffTariff  true  "date of staffTariff"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaffTariff(c *gin.Context) {
	fmt.Println("Method POST")
	var staffTariff models.CreateStaffTariff
	err := c.ShouldBind(&staffTariff)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.StaffTariff().CreateStaffTariff(c.Request.Context(), staffTariff)
	if err != nil {
		h.log.Error("error CreateStaffTariff", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateStaffTariff")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /staffTariff/{id} [get]
// @Summary      get staffTariff
// @Description  get staffTariff
// @Tags         staffTariffs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTariff"  Format(uuid)
// @Success      200  {object}   models.StaffTariff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaffTariff(c *gin.Context) {
	fmt.Println("Method GET")
	staffTarif := models.StaffTariff{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, staffTarif)
	if err != nil {
		fmt.Println("not found redis in cache")
	}
	if ok {
		c.JSON(http.StatusOK, staffTarif)
		return
	}

	resp, err := h.strg.StaffTariff().GetStaffTariff(c.Request.Context(), models.IdRequestStaffTariff{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error StaffTariff Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, 5*time.Minute)
	if err != nil {
		fmt.Println("error create redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staff/Tarif{id} [put]
// @Summary      updateda staffTarif
// @Description   api fot update staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTarif"
// @Param        staffTarif    body     models.CreateStaffTariff  true  "id of staffTarif"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffTariff(c *gin.Context) {
	fmt.Println("Method PUT")
	var staffTariff models.StaffTariff
	err := c.ShouldBindJSON(&staffTariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staffTariff.Id = c.Param("id")
	resp, err := h.strg.StaffTariff().UpdateStaffTariff(c.Request.Context(), staffTariff)
	if err != nil {
		fmt.Println("error StaffTariff Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Delete(c.Request.Context(), staffTariff.Id)
	if err != nil {
		fmt.Println("error delete redis", err.Error())
	}
}

// ListAccounts godoc
// @Router       /staffTariff/{id} [delete]
// @Summary      delete staffTariff
// @Description   api fot delete staffTariff
// @Tags         staffTariffs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTariff"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaffTariff(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error StaffTariff Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.StaffTariff().DeleteStaffTariff(c.Request.Context(), models.IdRequestStaffTariff{Id: id})
	if err != nil {
		h.log.Error("error StaffTariff Delete:", logger.Error(err))
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
// @Router       /staffTariff [get]
// @Summary      List staffTariff
// @Description  get staffTariff
// @Tags         staffTariffs
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllStaffTariff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaffTariff(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllStaffTariffs")
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

	resp, err := h.strg.StaffTariff().GetAllStaffTariff(c.Request.Context(), models.GetAllStaffTariffRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error StaffTariffs GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllStaffTariffs")
	c.JSON(http.StatusOK, resp)
}
