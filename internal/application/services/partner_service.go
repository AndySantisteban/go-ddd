package services

import (
	"github.con/AndyGo/go-ddd/internal/application/command"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
	"github.con/AndyGo/go-ddd/internal/domain/repositories"
)

type PartnerService struct {
	partnerRepository repositories.PLSPartnerRepository
}

func NewPartnerService(
	partnerRepository repositories.PLSPartnerRepository,
) *PartnerService {
	return &PartnerService{partnerRepository: partnerRepository}
}

func (s *PartnerService) GetAllPartners(cmd command.ListPartnerCommand) (*entities.DataSourceResponse[*entities.ValidatedPLSPartner], error) {

	datasourceRequest := entities.DataSourceRequest{
		Page:     cmd.DatasourceRequest.Page,
		PageSize: cmd.DatasourceRequest.PageSize,
		Sort:     cmd.DatasourceRequest.Sort,
		Filter:   cmd.DatasourceRequest.Filter,
	}
	return s.partnerRepository.GetAll(datasourceRequest)
}

func (s *PartnerService) FindPartnerByID(id string) (*entities.ValidatedPLSPartner, error) {
	return s.partnerRepository.FindByID(id)
}
