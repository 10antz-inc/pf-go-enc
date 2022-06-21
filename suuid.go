package enc

import (
	"fmt"

	guuid "github.com/google/uuid"
	"github.com/lithammer/shortuuid"
)

func SUUID() string {
	return shortuuid.New()
}

func UUIDtoSUUID(uuid string) (string, error) {
	if v, err := guuid.Parse(uuid); err != nil {
		return "", fmt.Errorf("failed to parse: %w", err)
	} else {
		return shortuuid.DefaultEncoder.Encode(v), nil
	}
}

func SUUIDtoUUID(suuid string) (string, error) {
	if v, err := shortuuid.DefaultEncoder.Decode(suuid); err != nil {
		return "", fmt.Errorf("failed to decode: %w", err)
	} else {
		return v.String(), nil
	}
}
