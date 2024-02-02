package entities

type ValidatedALLNote struct {
	ALLNote
	isValidated bool
}

func (van *ValidatedALLNote) IsValid() bool {
	return van.isValidated
}

func NewValidatedALLNote(note *ALLNote) (*ValidatedALLNote, error) {
	if err := note.validate(); err != nil {
		return nil, err
	}

	return &ValidatedALLNote{
		ALLNote:     *note,
		isValidated: true,
	}, nil
}
