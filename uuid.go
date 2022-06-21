package enc

import (
	"github.com/google/uuid"
)

func UUID() string {
	if v, err := uuid.NewRandom(); err != nil {
		return ""
	} else {
		return v.String()
	}
}
