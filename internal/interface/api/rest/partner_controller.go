package rest

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.con/AndyGo/go-ddd/internal/application/command"
	"github.con/AndyGo/go-ddd/internal/application/interfaces"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type PartnerController struct {
	service interfaces.PartnerService
}

func NewPartnerController(e *echo.Echo, service interfaces.PartnerService) *PartnerController {
	controller := &PartnerController{
		service: service,
	}
	e.GET("/partner", controller.GetAllPartners)
	e.GET("/partner/:id", controller.GetPartnersByID)

	return controller
}

func (pc *PartnerController) GetAllPartners(c echo.Context) error {
	datasourceRequest := c.QueryParam("DatasourceRequest")
	var dsRequest entities.DataSourceRequest
	if err := json.Unmarshal([]byte(datasourceRequest), &dsRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid DatasourceRequest JSON",
			"message": err.Error(),
		})
	}

	partners, err := pc.service.GetAllPartners(command.ListPartnerCommand{
		DatasourceRequest: dsRequest,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch partners",
		})
	}

	return c.JSON(http.StatusOK, partners)
}

func (pc *PartnerController) GetPartnersByID(c echo.Context) error {
	id := string(c.Param("id"))

	product, err := pc.service.FindPartnerByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to fetch partner",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}
