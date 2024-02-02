package mapper

import (
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/create"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

func NewNoteResultFromEntity(partner *entities.ValidatedPLSPartner) command.CreateALLNoteCommandResult {
	return command.CreateALLNoteCommandResult{}
}
