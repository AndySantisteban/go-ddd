package services

import (
	commandCreate "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/create"
	command "github.con/AndyGo/go-ddd/internal/application/command/ALLNote/query"
	"github.con/AndyGo/go-ddd/internal/domain/entities"
	"github.con/AndyGo/go-ddd/internal/domain/repositories"
)

type ALLNoteService struct {
	allNoteRepository repositories.ALLNoteRepository
}

func NewALLNoteService(
	allNoteRepository repositories.ALLNoteRepository,
) *ALLNoteService {
	return &ALLNoteService{allNoteRepository: allNoteRepository}
}

func (s *ALLNoteService) GetAllALLNote(cmd command.ListALLNoteCommand) (*entities.DataSourceResponse[*entities.ValidatedALLNote], error) {

	datasourceRequest := entities.DataSourceRequest{
		Page:     cmd.DatasourceRequest.Page,
		PageSize: cmd.DatasourceRequest.PageSize,
		Sort:     cmd.DatasourceRequest.Sort,
		Filter:   cmd.DatasourceRequest.Filter,
	}
	return s.allNoteRepository.GetAll(datasourceRequest)
}

func (s *ALLNoteService) FindALLNoteByID(id string) (*entities.ValidatedALLNote, error) {
	return s.allNoteRepository.FindByID(id)
}

func (s *ALLNoteService) Create(cmd commandCreate.CreateALLNoteCommand) (commandCreate.CreateALLNoteCommandResult, error) {

	newNoteID, err := s.allNoteRepository.Create(cmd.ALLNote)
	if err != nil {
		return commandCreate.CreateALLNoteCommandResult{}, err
	}

	result := commandCreate.CreateALLNoteCommandResult{
		Result: newNoteID,
	}

	return result, nil

}
