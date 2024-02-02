package entities

import (
	"errors"
	"time"
)

type ALLNote struct {
	Uid                 int
	ParentUid           *string
	UserUid             string
	SubjectOld          *string
	CategoryUid         int
	SubCategory         int
	Note                string
	Note_Plain          string
	Priority            int
	CommunicationCode   *int
	IsReminder          bool
	CompletedBy         *string
	CompletedDate       *time.Time
	ScheduledStart      *time.Time
	ScheduledEnd        *time.Time
	OldRecId            *string
	AppTimeStamp        *time.Time
	AppLastUpdatedBy    *string
	AppCreatedBy        *string
	ActivityUid         *string
	Subject             *string
	ContactName         *string
	AppTimeStampYear    int
	SysTimeStamp        []uint8
	Note_PlainResume    *string
	ParentDepartmentUid int
	ParentAccount       string
	ParentType          *int
	AppLastUpdated      *time.Time
}

func (p *ALLNote) validate() error {
	if p.Uid <= 0 {
		return errors.New("invalid Uid")
	}

	// if p.ParentUid == "" {
	// 	return errors.New("invalid Parent Uid")
	// }

	// if p.UserUid == "" {
	// 	return errors.New("invalid User Uid")
	// }

	// if p.CategoryUid == 0 {
	// 	return errors.New("invalid Category Uid")
	// }

	return nil
}