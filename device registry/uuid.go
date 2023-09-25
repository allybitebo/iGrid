package registry

import (
	"github.com/gofrs/uuid"
	"github.com/piusalfred/registry/pkg/errors"
)

// UUIDProvider specifies an API for generating unique identifiers.
type UUIDProvider interface {
	// ID generates the unique identifier.
	ID() (string, error)
}

// ErrGeneratingID indicates errors in generating UUID
var ErrGeneratingID = errors.New("generating id failed")

var _ UUIDProvider = (*uuidProvider)(nil)

type uuidProvider struct{}

// New instantiates a UUID provider.
func New() UUIDProvider {

	return &uuidProvider{}
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(ErrGeneratingID, err)
	}

	return id.String(), nil
}
