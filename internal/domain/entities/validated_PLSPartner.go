package entities

type ValidatedPLSPartner struct {
	PLSPartner
	isValidated bool
}

func (vp *ValidatedPLSPartner) IsValid() bool {
	return vp.isValidated
}

func NewValidatedPLSPartner(partner *PLSPartner) (*ValidatedPLSPartner, error) {
	if err := partner.validate(); err != nil {
		return nil, err
	}

	return &ValidatedPLSPartner{
		PLSPartner:  *partner,
		isValidated: true,
	}, nil
}
