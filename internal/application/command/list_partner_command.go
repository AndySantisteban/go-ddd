package command

import (
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ListPartnerCommand struct {
	DatasourceRequest entities.DataSourceRequest
}
