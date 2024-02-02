package repositories

import (
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ALLNoteRepository interface {
	FindByID(id string) (*entities.ValidatedALLNote, error)
	Create(note entities.ALLNote) (int, error)
	GetAll(datasourceRequest entities.DataSourceRequest) (*entities.DataSourceResponse[*entities.ValidatedALLNote], error)
}
