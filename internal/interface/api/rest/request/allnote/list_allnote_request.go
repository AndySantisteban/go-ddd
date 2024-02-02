package request

import (
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ListALLNoteRequest struct {
	DataSourceRequest entities.DataSourceRequest `json:"dataSourceRequest"`
}

func (req *ListALLNoteRequest) ToCreatePartnerCommand() (*command.ListALLNoteCommand, error) {

	return &command.ListALLNoteCommand{
		DatasourceRequest: req.DataSourceRequest,
	}, nil
}
