package model

import (
	"fmt"
	"io"
	"time"
)

type Date time.Time

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (d *Date) UnmarshalGQL(v interface{}) error {
	datetime, ok := v.(string)
	if !ok {
		return fmt.Errorf("Date must be a string")
	}
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (d Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(time.Time(d).String()))
}
