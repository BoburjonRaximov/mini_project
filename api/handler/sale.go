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
// @Router       /sale [post]
// @Summary      create sale
// @Description  api for create sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale    body     models.CreateSale  true  "date of sale"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateSale(c *gin.Context) {
	fmt.Println("Method POST")
	var sale models.CreateSale
	err := c.ShouldBind(&sale)
	if err != nil {
		h.log.Error("error ShouldBind", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.strg.Sale().CreateSale(sale)
	if err != nil {
		h.log.Error("error CreateSale", logger.Error(err))
		c.JSON(http.StatusInternalServerError, " error CreateSale")
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListAccounts godoc
// @Router       /sale/{id} [get]
// @Summary      get sale
// @Description  get sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of s"  Format(uuid)
// @Success      200  {object}   models.Sale
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetSale(c *gin.Context) {
	fmt.Println("Method GET")
	id := c.Param("id")

	resp, err := h.strg.Sale().GetSale(models.IdRequestSale{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Sale Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /sale/{id} [put]
// @Summary      updateda sale
// @Description   api fot update sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of sale"
// @Param        sale    body     models.CreateSale  true  "id of sale"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateSale(c *gin.Context) {
	fmt.Println("Method PUT")
	var sale models.Sale
	err := c.ShouldBind(&sale)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	sale.Id = c.Param("id")
	resp, err := h.strg.Sale().UpdateSale(sale)
	if err != nil {
		fmt.Println("error Sale Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /sale/{id} [delete]
// @Summary      delete sale
// @Description   api fot delete sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of sale"
// @Success      200  {strig}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteSale(c *gin.Context) {
	fmt.Println("Method DELETE")
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Sale Delete:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Sale().DeleteSale(models.IdRequestSale{Id: id})
	if err != nil {
		h.log.Error("error Sale Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccounts godoc
// @Router       /sale [get]
// @Summary      List sales
// @Description  get sale
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {object}  models.GetAllSale
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllSale(c *gin.Context) {
	fmt.Println("Method GetAll")
	h.log.Info("request GetAllSales")
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

	resp, err := h.strg.Sale().GetAllSale(models.GetAllSaleRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Sale GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllSale")
	c.JSON(http.StatusOK, resp)
}
