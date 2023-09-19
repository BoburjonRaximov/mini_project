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
// @Router       /staffTarif [post]
// @Summary      create staffTarif
// @Description  api for create staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        staffTarif    body     models.CreateStaffTarif  true  "date of staffTarif"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaffTarif(c *gin.Context) {
	fmt.Println("Method POST")
	var staffTarif models.CreateStaffTarif
	err := c.ShouldBind(&staffTarif)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.StaffTarif().CreateStaffTarif(staffTarif)
	if err != nil {
		h.log.Error("error CreateStaff", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateStaffTarif")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /staffTarif/{id} [get]
// @Summary      get staffTarif
// @Description  get staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTarif"  Format(uuid)
// @Success      200  {object}   models.StaffTarif
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaffTarif(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequestStaffTarif{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error StaffTarif Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /staff/Tarif{id} [put]
// @Summary      updateda staffTarif
// @Description   api fot update staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTarif"
// @Param        staffTarif    body     models.CreateStaffTarif  true  "id of staffTarif"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffTarif(c *gin.Context) {
	fmt.Println("Method PUT")
	var staffTarif models.StaffTarif
	err := c.ShouldBind(&staffTarif)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staffTarif.Id = c.Param("id")
	resp, err := h.strg.StaffTarif().UpdateStaffTarif(staffTarif)
	if err != nil {
		fmt.Println("error StaffTarif Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /staffTarif/{id} [delete]
// @Summary      delete staffTarif
// @Description   api fot delete staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staffTarif"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error StaffTarif Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.StaffTarif().DeleteStaffTarif(models.IdRequestStaffTarif{Id: id})
	if err != nil {
		h.log.Error("error StaffTarif Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /staffTarif [get]
// @Summary      List staffTarif
// @Description  get staffTarif
// @Tags         staffTarifs
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllStaffTarif
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaffTarif(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllStaffTarifs")
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

	resp, err := h.strg.StaffTarif().GetAllStaffTarif(models.GetAllStaffTarifRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error StaffTarifs GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllStaffTarifs")
	c.JSON(http.StatusOK, resp)
}
