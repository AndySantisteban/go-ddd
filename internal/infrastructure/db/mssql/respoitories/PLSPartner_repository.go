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

type SQLPartnerRepository struct {
}

func NewPLSPartnerRepository() repositories.PLSPartnerRepository {
	return &SQLPartnerRepository{}
}

func (repo *SQLPartnerRepository) FindByID(id string) (*entities.ValidatedPLSPartner, error) {
	sqlserver, err := mssql.NewConnection("")

	if err != nil {
		return nil, err
	}
	var partners []models.PLSPartner
	Uid := fmt.Sprintf("Uid = '%s'", id)
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("Uid", "Account").
		From("PLSPartner").
		Where(Uid)

	rows, _, err := query.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var partner models.PLSPartner

		if err := rows.Scan(&partner.Uid, &partner.Account); err != nil {
			log.Fatal("Error fetching or PLSPartner Uid mismatch", err)
			return nil, err

		}
		partners = append(partners, partner)
	}
	sqlserver.Close()

	fmt.Sprintln(" partners: ", partners)

	return mappers.FromDBLSPartner(&partners[0])
}

func (repo *SQLPartnerRepository) GetAll(datasourceRequest entities.DataSourceRequest) (*entities.DataSourceResponse[*entities.ValidatedPLSPartner], error) {
	sqlserver, err := mssql.NewConnection("")

	if err != nil {
		panic("Failed to connect to database")
	}
	var partners []models.PLSPartner
	query := mssql.NewSQLQueryBuilder(sqlserver).
		Select("Uid", "Account", "IsCompany", "Company").
		From("PLSPartner").
		ApplyDataSourceRequest(&datasourceRequest)

	rows, total, err := query.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var partner models.PLSPartner

		if err := rows.Scan(&partner.Uid, &partner.Account, &partner.IsCompany, &partner.Company); err != nil {
			log.Fatal("Error fetching or PLSPartner Uid mismatch", err)

		}
		partners = append(partners, partner)
	}
	sqlserver.Close()
	partnersList := make([]*entities.ValidatedPLSPartner, len(partners))
	for i, dbPLSPartner := range partners {
		partnersList[i], err = mappers.FromDBLSPartner(&dbPLSPartner)
		if err != nil {
			return nil, err
		}
	}

	response := entities.DataSourceResponse[*entities.ValidatedPLSPartner]{
		Data:  partnersList,
		Total: total,
	}

	return &response, nil
}
