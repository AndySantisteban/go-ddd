package sqlite_test

import (
	"database/sql"
	"log"
	"testing"

	"github.con/AndyGo/go-ddd/internal/domain/entities"
	"github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql"
	models "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/models"
)

func setupDatabase() (*sql.DB, error) {
	database, err := mssql.NewConnection("CenturionNotes")
	if err != nil {
		panic("Failed to connect to database")
	}
	return database, err

}

func TestSQLPLSPartnerRepository_Select(t *testing.T) {
	request := entities.DataSourceRequest{
		Page:     1,
		PageSize: 10000,
	}

	t.Log("request ", request)

	sqlserver, err := setupDatabase()

	if err != nil {
		t.Error("Failed to connect to database")
	}
	notes := make([]models.ALLNote, request.PageSize)

	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("*").
		From("ALLNote").
		ApplyDataSourceRequest(&request)

	rows, total, err := query.Query()

	if err != nil {
		t.Error("Error fetching or PLSPartner mismatch", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var note models.ALLNote
		index := 0
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
		notes[index] = note
		index++
	}
	t.Log("partners length:", total)
}

func TestSQLPLSPartnerRepository_SelectByID(t *testing.T) {

	sqlserver, err := setupDatabase()

	if err != nil {
		t.Error("Failed to connect to database")
	}
	var partners []models.PLSPartner
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("Uid", "Account").
		From("PLSPartner").
		Where("Uid = '001696c6a395421189eb14cf85b881df'")

	rows, _, err := query.Query()

	if err != nil {
		t.Error("Error fetching or PLSPartner mismatch", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var partner models.PLSPartner

		if err := rows.Scan(&partner.Uid, &partner.Account); err != nil {
			t.Error("Error fetching or PLSPartner Uid mismatch", err)
			return
		}
		partners = append(partners, partner)
		t.Logf("partners %p \n", partners)
	}

}
