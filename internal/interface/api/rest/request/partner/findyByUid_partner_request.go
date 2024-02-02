package request

import (
	"github.con/AndyGo/go-ddd/internal/application/command"
)

type FindByUidPartnerRequest struct {
	Uid string `json:"Uid"`
}

func (req *FindByUidPartnerRequest) ToCreatePartnerCommand() (*command.FinByUidPartnerCommand, error) {

	return &command.FinByUidPartnerCommand{
		Uid: req.Uid,
	}, nil
}
