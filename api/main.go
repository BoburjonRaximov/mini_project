package api

import (
	"new_project/api/docs"
	"new_project/api/handler"
	"new_project/pkg/helper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func NewServer(h *handler.Handler) *gin.Engine {

	r := gin.Default()
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "First API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Use(helper.StartMiddleware)
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch", h.GetAllBranch, helper.Logger)
	r.GET("/branch/:id", h.GetBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.Use(helper.Logger)
	r.POST("/staff", h.CreateStaff)
	r.GET("/staff", h.GetAllStaff, helper.AuthMiddleware)
	r.GET("/staff/:id", h.GetStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	r.POST("/staffTariff", h.CreateStaffTariff)
	r.GET("/staffTariff", h.GetAllStaffTariff)
	r.GET("/staffTariff/:id", h.GetStaffTariff)
	r.PUT("/staffTariff/:id", h.UpdateStaffTariff)
	r.DELETE("/staffTariff/:id", h.DeleteStaffTariff)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale", h.GetAllSale)
	r.GET("/sale/:id", h.GetSale)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	r.POST("/staffTransaction", h.CreateStaffTransaction)
	r.GET("/staffTransaction", h.GetAllStaffTransaction)
	r.GET("/staffTransaction/:id", h.GetStaffTransaction)
	r.PUT("/staffTransaction/:id", h.UpdateStaffTransaction)
	r.DELETE("/staffTransaction/:id", h.DeleteStaffTransaction)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
