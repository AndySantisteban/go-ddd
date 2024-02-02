package entities

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewALLNote(t *testing.T) {
	newUid := uuid.NewString()
	note := &ALLNote{
		Uid:         1,
		ParentUid:   &newUid,
		CategoryUid: 1,
		UserUid:     uuid.NewString(),
	}

	if note.Uid == 0 {
		t.Errorf("Expected All Note uid '%s', but got %o", newUid, note.Uid)
	}

	if note.ParentUid == &newUid {
		t.Errorf("Expected All Note ParentUid '%s', but got nil", *note.ParentUid)
	}

	if note.CategoryUid == 0 {
		t.Errorf("Expected All Note CategoryUid '%o', but got nil", note.CategoryUid)
	}

}
