package entities

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPLSPartner(t *testing.T) {
	newUid := uuid.NewString()
	partner := &PLSPartner{
		Uid: newUid,
	}

	if partner.Uid != newUid {
		t.Errorf("Expected partner uid '%s', but got %s", newUid, partner.Uid)
	}

}
