package converter

import (
	"testing"
)

func TestNewMappingFromString(t *testing.T) {
	convert := "from:user,to:other"

	mapping, err := NewMappingFromString(convert)
	if err != nil {
		t.Fatal()
	}

	if mapping.from != "user" && mapping.to != "other" {
		t.Fatal()
	}
}
