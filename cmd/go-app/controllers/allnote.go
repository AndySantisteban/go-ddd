package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

// App struct
type AllNoteController struct {
	ctx context.Context
}

const BaseUrl = "http://localhost:8080/allnote"

// NewApp creates a new App application struct
func NewControllerAllNote() *AllNoteController {
	return &AllNoteController{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *AllNoteController) Startup(ctx context.Context) {
	a.ctx = ctx
}

type ALLNoteDTO struct {
	Uid                 int        `json:"Data"`
	ParentUid           *string    `json:"ParentUid"`
	UserUid             string     `json:"UserUid"`
	SubjectOld          *string    `json:"SubjectOld"`
	CategoryUid         int        `json:"CategoryUid"`
	SubCategory         int        `json:"SubCategory"`
	Note                string     `json:"Note"`
	Note_Plain          string     `json:"Note_Plain"`
	Priority            int        `json:"Priority"`
	CommunicationCode   *int       `json:"CommunicationCode"`
	IsReminder          bool       `json:"IsReminder"`
	CompletedBy         *string    `json:"CompletedBy"`
	CompletedDate       *time.Time `json:"CompletedDate"`
	ScheduledStart      *time.Time `json:"ScheduledStart"`
	ScheduledEnd        *time.Time `json:"ScheduledEnd"`
	OldRecId            *string    `json:"OldRecId"`
	AppTimeStamp        *time.Time `json:"AppTimeStamp"`
	AppLastUpdatedBy    *string    `json:"AppLastUpdatedBy"`
	AppCreatedBy        *string    `json:"AppCreatedBy"`
	ActivityUid         *string    `json:"ActivityUid"`
	Subject             *string    `json:"Subject"`
	ContactName         *string    `json:"ContactName"`
	AppTimeStampYear    int        `json:"AppTimeStampYear"`
	SysTimeStamp        []uint8    `json:"SysTimeStamp"`
	Note_PlainResume    *string    `json:"Note_PlainResume"`
	ParentDepartmentUid int        `json:"ParentDepartmentUid"`
	ParentAccount       string     `json:"ParentAccount"`
	ParentType          *int       `json:"ParentType"`
	AppLastUpdated      *time.Time `json:"AppLastUpdated"`
}

// Greet returns a greeting for the given name
func (a *AllNoteController) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Greet returns a greeting for the given name
func (a *AllNoteController) DomainExpansion(name string) string {
	return fmt.Sprintf("Ryōiki Tenkai : %s", name)
}

func (a *AllNoteController) GreetNight(name string) string {
	return fmt.Sprintf("Bye %s !", name)
}

type DataSourceResponseDTO struct {
	Data  []*ALLNoteDTO `json:"Data"`
	Total int           `json:"Total"`
}

func (a *AllNoteController) GetData(datasourceRequest *entities.DataSourceRequest) (*DataSourceResponseDTO, error) {
	baseUrl, err := url.Parse(BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal DataSourceRequest to JSON
	requestData, err := json.Marshal(datasourceRequest)
	if err != nil {
		return nil, err
	}

	values := baseUrl.Query()
	values.Add("DatasourceRequest", string(requestData))
	baseUrl.RawQuery = values.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Deserialize the JSON response
	var dataSourceResponse DataSourceResponseDTO
	if err := json.Unmarshal(body, &dataSourceResponse); err != nil {
		return nil, err
	}

	return &dataSourceResponse, nil
}
