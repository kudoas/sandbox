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

var _ graphql.Marshaler = (*URL)(nil)

func (u URL) MarshalGQL(w io.Writer) { io.WriteString(w, fmt.Sprintf(`"%s"`, u.URL.String())) }

var _ graphql.Unmarshaler = (*URL)(nil)

func (u *URL) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		result, err := url.Parse(v)
		if err != nil {
			return err
		}
		u = &URL{*result}
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
