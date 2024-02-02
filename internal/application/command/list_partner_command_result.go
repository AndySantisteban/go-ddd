package command

import (
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type ListPartnerCommandResult struct {
	Result *entities.DataSourceResponse[entities.PLSPartner]
}
