package sqlite_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.con/AndyGo/go-ddd/internal/domain/entities"
	"github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql"
	models "github.con/AndyGo/go-ddd/internal/infrastructure/db/mssql/models"
)

func setupDatabase() (*sql.DB, error) {
	database, err := mssql.NewConnection("")
	if err != nil {
		panic("Failed to connect to database")
	}
	return database, err

}

func TestSQLPLSPartnerRepository_Select(t *testing.T) {
	request := entities.DataSourceRequest{
		Page:     1,
		PageSize: 20,
	}

	t.Log("request ", request)

	sqlserver, err := setupDatabase()

	if err != nil {
		t.Error("Failed to connect to database")
	}
	var partners []models.PLSPartner
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("Uid", "Account").
		From("PLSPartner").
		ApplyDataSourceRequest(&request)

	rows, total, err := query.Query()

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
	}
	t.Log("partners length:", total)
	for i := 0; i < len(partners); i++ {
		fmt.Printf("Partner :  account: %s , Uid: %s \n", partners[i].Account, partners[i].Uid)

	}

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
