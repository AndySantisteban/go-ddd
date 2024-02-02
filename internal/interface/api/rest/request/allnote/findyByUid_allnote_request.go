package request

import (
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
)

type FindByUidALLNoteRequest struct {
	Uid string `json:"Uid"`
}

func (req *FindByUidALLNoteRequest) ToCreateALLNoteCommand() (*command.FinByUidALLNoteCommand, error) {

	return &command.FinByUidALLNoteCommand{
		Uid: req.Uid,
	}, nil
}
