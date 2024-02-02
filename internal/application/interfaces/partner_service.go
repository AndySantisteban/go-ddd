package interfaces

import (
	command "github.con/AndyGo/go-ddd/internal/application/command"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type PartnerService interface {
	GetAllPartners(listPartnerCommand command.ListPartnerCommand) (*entities.DataSourceResponse[*entities.ValidatedPLSPartner], error)
	FindPartnerByID(id string) (*entities.ValidatedPLSPartner, error)
}
