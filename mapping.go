package converter

import (
	"errors"
	"strings"
)

type mapping struct {
	from string
	to   string
}

var ErrInvalidateFormat = errors.New("invalid format")

func newMappingFromString(convert string) (*mapping, error) {
	chunks := strings.SplitN(convert, ",", 2)
	if len(chunks) != 2 {
		return nil, ErrInvalidateFormat
	}
	if strings.HasPrefix(chunks[0], "to:") { // 排个序
		chunks[0], chunks[1] = chunks[1], chunks[0]
	}

	from := strings.TrimPrefix(chunks[0], "from:")
	from = strings.Trim(from, " ")
	to := strings.TrimPrefix(chunks[1], "to:")
	to = strings.Trim(to, " ")

	return &mapping{from, to}, nil
}
