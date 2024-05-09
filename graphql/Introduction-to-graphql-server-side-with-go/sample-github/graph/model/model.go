package model

import (
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

type URL struct {
	url.URL
}

var (
	_ graphql.Marshaler   = (*URL)(nil)
	_ graphql.Unmarshaler = (*URL)(nil)
)

// MarshalGQL implements the graphql.Marshaler interface
func (u URL) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.URL.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *URL) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		if result, err := url.Parse(v); err != nil {
			return err
		} else {
			u = &URL{*result}
		}
		return nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return err
		}
		u = &URL{*result}
		return nil
	default:
		return fmt.Errorf("%T is not a url.URL", v)
	}
}
