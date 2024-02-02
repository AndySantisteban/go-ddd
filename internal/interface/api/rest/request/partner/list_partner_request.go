package request

import (
	"github.con/AndyGo/go-ddd/internal/application/command"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ListPartnerRequest struct {
	DataSourceRequest entities.DataSourceRequest `json:"dataSourceRequest"`
}

func (req *ListPartnerRequest) ToCreatePartnerCommand() (*command.ListPartnerCommand, error) {

	return &command.ListPartnerCommand{
		DatasourceRequest: req.DataSourceRequest,
	}, nil
}
