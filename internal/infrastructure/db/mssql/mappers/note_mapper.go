package mssql

import (
	"github.com/devfeel/mapper"
	entities "github.con/AndyGo/go-ddd/internal/domain/entities"
	models "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/models"
)

func ToDBALLNote(note *entities.ValidatedALLNote) *models.ALLNote {
	ALLNote := &models.ALLNote{}
	mapper.AutoMapper(note, ALLNote)
	return ALLNote
}

func FromDBALLNote(dbALLNote *models.ALLNote) (*entities.ValidatedALLNote, error) {
	ALLNote := &entities.ALLNote{}
	mapper.AutoMapper(dbALLNote, ALLNote)

	return entities.NewValidatedALLNote(ALLNote)
}
