package command

import (
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ListALLNoteCommandResult struct {
	Result *entities.DataSourceResponse[entities.ALLNote]
}
