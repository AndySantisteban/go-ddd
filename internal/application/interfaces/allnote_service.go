package interfaces

import (
	commandCreate "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/create"
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ALLNoteService interface {
	Create(ALLNote commandCreate.CreateALLNoteCommand) (commandCreate.CreateALLNoteCommandResult, error)
	GetAllALLNote(listALLNoteCommand command.ListALLNoteCommand) (*entities.DataSourceResponse[*entities.ValidatedALLNote], error)
	FindALLNoteByID(id string) (*entities.ValidatedALLNote, error)
}
