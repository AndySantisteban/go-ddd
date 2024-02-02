package rest

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	commandCreate "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/create"
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
	"github.con/AndyGo/go-ddd/internal/application/interfaces"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ALLNoteController struct {
	service interfaces.ALLNoteService
}

func NewALLNoteController(e *echo.Echo, service interfaces.ALLNoteService) *ALLNoteController {
	controller := &ALLNoteController{
		service: service,
	}
	e.GET("/allnote", controller.GetAll)
	e.GET("/allnote/:id", controller.GetByID)
	e.POST("/allnote", controller.Create)

	return controller
}

func (pc *ALLNoteController) GetAll(c echo.Context) error {
	datasourceRequest := c.QueryParam("DatasourceRequest")
	var dsRequest entities.DataSourceRequest
	if err := json.Unmarshal([]byte(datasourceRequest), &dsRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid DatasourceRequest JSON",
			"message": err.Error(),
		})
	}

	partners, err := pc.service.GetAllALLNote(command.ListALLNoteCommand{
		DatasourceRequest: dsRequest,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to fetch notes",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, partners)
}

func (pc *ALLNoteController) GetByID(c echo.Context) error {
	id := string(c.Param("id"))

	product, err := pc.service.FindALLNoteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to fetch note",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ALLNoteController) Create(c echo.Context) error {
	// id := string(c.Param("id"))
	var json_map entities.ALLNote
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to create note",
			"message": err.Error(),
		})
	}

	product, err := pc.service.Create(commandCreate.CreateALLNoteCommand{
		ALLNote: json_map,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to fetch note",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}
