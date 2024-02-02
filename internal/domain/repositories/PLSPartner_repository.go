package repositories

import (
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

type PLSPartnerRepository interface {
	FindByID(id string) (*entities.ValidatedPLSPartner, error)
	GetAll(datasourceRequest entities.DataSourceRequest) (*entities.DataSourceResponse[*entities.ValidatedPLSPartner], error)
}
