package converter

import (
	"testing"
)

func TestNewMappingFromString(t *testing.T) {
	converts := []string{"from:user,to:other", "from:user,  to:other", "from:  user, to: other"}

	for _, convert := range converts {
		mapping, err := NewMappingFromString(convert)
		if err != nil {
			t.Fatal()
		}

		if mapping.from != "user" && mapping.to != "other" {
			t.Fatal()
		}
	}
}
