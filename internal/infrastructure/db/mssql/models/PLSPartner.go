package mssql

import (
	"time"
)

type PLSPartner struct {
	Uid                            string
	Account                        string
	DepartmentUid                  int
	IsCompany                      bool
	Company                        string
	ContacName                     string
	Salutation                     string
	FirstName                      string
	MiddleName                     string
	LastName                       string
	FullName                       string
	Street                         string
	City                           string
	State                          string
	ZipCode                        string
	HomePhone                      string
	WorkPhone                      string
	MobilePhone                    string
	Fax                            string
	TIN                            string
	TINType                        int
	TINMask                        string
	ERISA                          bool
	Payee                          string
	UsePayee                       bool
	AccountType                    int
	BirthDay                       time.Time
	RolodexChanged                 bool
	PrintStatementFor              int
	Categories                     string
	Email                          string
	ACHSendDepositNotificationFlag int
	ACHReceivingDFI                string
	ACHAccountNumber               string
	ACHIndividualIdNumber          string
	ACHIndividualName              string
	ACHBankAccountType             int
	ACHStatus                      int
	AppTimeStamp                   time.Time
	AppCreatedBy                   string
	AppCreationDate                time.Time
	AppLastUpdatedBy               string
	SysTimeStamp                   [5]byte
	EmailFormat                    int
	DeliveryOptions                int
	Module                         int
	TINEncrypted                   string
	ACHAddendaInfo                 string
	ACHAddendaEnable               bool
}
