package mssql

import (
	"fmt"
	"log"

	"github.con/AndyGo/go-ddd/internal/domain/entities"
	"github.con/AndyGo/go-ddd/internal/domain/repositories"
	"github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql"
	mappers "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/mappers"
	models "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/models"
)

type SQLALLNoteRepository struct {
}

func NewALLNoteRepository() repositories.ALLNoteRepository {
	return &SQLALLNoteRepository{}
}

// Create implements repositories.ALLNoteRepository.
func (*SQLALLNoteRepository) Create(note entities.ALLNote) (int, error) {
	sqlserver, err := mssql.NewConnection("CenturionNotes")

	if err != nil {
		return 0, err
	}

	queryBuilder := mssql.NewSQLQueryBuilder(sqlserver)

	insertedNote, err := queryBuilder.Insert("ALLNote", note, "Uid", "AppTimeStamp", "AppLastUpdated", "SysTimeStamp")
	if err != nil {
		return 0, err
	}
	sqlserver.Close()
	// fmt.Println("Note inserted:", insertedNote)
	noteconvert, ok := insertedNote.(*entities.ALLNote)
	if !ok {
		return 0, err
	}

	return noteconvert.Uid, nil

}

// FindByID implements repositories.ALLNoteRepository.
func (*SQLALLNoteRepository) FindByID(id string) (*entities.ValidatedALLNote, error) {
	sqlserver, err := mssql.NewConnection("CenturionNotes")

	if err != nil {
		return nil, err
	}
	var notes []models.ALLNote
	Uid := fmt.Sprintf("Uid = '%s'", id)
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("*").
		From("ALLNote").
		Where(Uid)

	rows, _, err := query.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.ALLNote

		if err := rows.Scan(&note.Uid,
			&note.ParentUid,
			&note.UserUid,
			&note.SubjectOld,
			&note.CategoryUid,
			&note.SubCategory,
			&note.Note,
			&note.Note_Plain,
			&note.Priority,
			&note.CommunicationCode,
			&note.IsReminder,
			&note.CompletedBy,
			&note.CompletedDate,
			&note.ScheduledStart,
			&note.ScheduledEnd,
			&note.OldRecId,
			&note.AppTimeStamp,
			&note.AppLastUpdatedBy,
			&note.AppCreatedBy,
			&note.ActivityUid,
			&note.Subject,
			&note.ContactName,
			&note.AppTimeStampYear,
			&note.SysTimeStamp,
			&note.Note_PlainResume,
			&note.ParentDepartmentUid,
			&note.ParentAccount,
			&note.ParentType,
			&note.AppLastUpdated); err != nil {
			log.Fatal("Error fetching or Note Uid mismatch", err)
			return nil, err

		}
		notes = append(notes, note)
	}
	sqlserver.Close()
	fmt.Sprintln(" notes: ", notes)

	return mappers.FromDBALLNote(&notes[0])
}

// GetAll implements repositories.ALLNoteRepository.
func (*SQLALLNoteRepository) GetAll(datasourceRequest entities.DataSourceRequest) (*entities.DataSourceResponse[*entities.ValidatedALLNote], error) {
	sqlserver, err := mssql.NewConnection("CenturionNotes")

	if err != nil {
		panic("Failed to connect to database")
	}
	var notes []models.ALLNote
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("*").
		From("ALLNote").
		ApplyDataSourceRequest(&datasourceRequest)

	rows, total, err := query.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var note models.ALLNote

		if err := rows.Scan(
			&note.Uid,
			&note.ParentUid,
			&note.UserUid,
			&note.SubjectOld,
			&note.CategoryUid,
			&note.SubCategory,
			&note.Note,
			&note.Note_Plain,
			&note.Priority,
			&note.CommunicationCode,
			&note.IsReminder,
			&note.CompletedBy,
			&note.CompletedDate,
			&note.ScheduledStart,
			&note.ScheduledEnd,
			&note.OldRecId,
			&note.AppTimeStamp,
			&note.AppLastUpdatedBy,
			&note.AppCreatedBy,
			&note.ActivityUid,
			&note.Subject,
			&note.ContactName,
			&note.AppTimeStampYear,
			&note.SysTimeStamp,
			&note.Note_PlainResume,
			&note.ParentDepartmentUid,
			&note.ParentAccount,
			&note.ParentType,
			&note.AppLastUpdated,
		); err != nil {
			log.Fatal("Error fetching or Note Uid mismatch", err)

		}
		sqlserver.Close()
		notes = append(notes, note)
	}

	notesList := make([]*entities.ValidatedALLNote, len(notes))
	for i, dbNotes := range notes {
		notesList[i], err = mappers.FromDBALLNote(&dbNotes)
		if err != nil {
			return nil, err
		}
	}

	response := entities.DataSourceResponse[*entities.ValidatedALLNote]{
		Data:  notesList,
		Total: total,
	}

	return &response, nil
}
