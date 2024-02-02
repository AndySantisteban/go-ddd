package mssql

import (
	"github.com/devfeel/mapper"
	entities "github.con/AndyGo/go-ddd/internal/domain/entities"
	models "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/models"
)

func ToDBPLSPartner(partner *entities.ValidatedPLSPartner) *models.PLSPartner {

	partnerEntity := &models.PLSPartner{}
	mapper.AutoMapper(partner, partnerEntity)

	return partnerEntity
}

func FromDBLSPartner(dbPLSPartner *models.PLSPartner) (*entities.ValidatedPLSPartner, error) {
	partner := &entities.PLSPartner{}
	mapper.AutoMapper(dbPLSPartner, partner)

	return entities.NewValidatedPLSPartner(partner)

}
