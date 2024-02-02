package mapper

import (
	"github.con/AndyGo/go-ddd/internal/application/command"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

func NewPartnerResultFromEntity(partner *entities.ValidatedPLSPartner) command.PartnerResult {
	return command.PartnerResult{
		PLSPartner: partner.PLSPartner,
	}
}
