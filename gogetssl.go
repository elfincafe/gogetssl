package gogetssl

import (
	"bytes"
	"errors"
	"time"
)

type (
	Api struct {
		LiveAPI string
		Key     string
	}
	Error struct {
		Error       bool   `json:"error"`
		Message     string `json:"message"`
		Description string `json:"description"`
	}
	DateTime struct {
		time.Time
	}
)

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, "\"'`"))
	formats := []string{"2006-01-02", "2006-01-02 15:04:05", "2006-01-02 15:04", time.RFC3339}
	for _, format := range formats {
		t, err := time.Parse(format, s)
		if err == nil {
			dt.Time = t
			return nil
		}
	}
	return errors.New("needs time formated value")
}
