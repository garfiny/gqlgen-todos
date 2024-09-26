package model

import (
	"fmt"
	"io"
)

//
// Most common scalars
//

type YesNo bool

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (y *YesNo) UnmarshalGQL(v interface{}) error {
	yes, ok := v.(string)
	if !ok {
		return fmt.Errorf("YesNo must be a string")
	}

	if yes == "yes" {
		*y = true
	} else {
		*y = false
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (y YesNo) MarshalGQL(w io.Writer) {
	if y {
		w.Write([]byte(`"yes"`))
	} else {
		w.Write([]byte(`"no"`))
	}
}

//
// Scalars that need access to the request context
//

// type Length float64

// // UnmarshalGQLContext implements the graphql.ContextUnmarshaler interface
// func (l *Length) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
// 	s, ok := v.(string)
// 	if !ok {
// 		return fmt.Errorf("Length must be a string")
// 	}
// 	length, err := ParseLength(s)
// 	if err != nil {
// 		return err
// 	}
// 	*l = length
// 	return nil
// }

// // MarshalGQLContext implements the graphql.ContextMarshaler interface
// func (l Length) MarshalGQLContext(ctx context.Context, w io.Writer) error {
// 	s, err := l.FormatContext(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	w.Write([]byte(strconv.Quote(s)))
// 	return nil
// }

// // ParseLength parses a length measurement string with unit on the end (eg: "12.45in")
// func ParseLength(string) (Length, error)

// // ParseLength formats the string using a value in the context to specify format
// func (l Length) FormatContext(ctx context.Context) (string, error)
