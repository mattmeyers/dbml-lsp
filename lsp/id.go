package lsp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ID is a request ID. This ID can be either a string or an int. All values
// are converted to a string when unmarshalled. The ID is converted back to an
// int when marshalled if the IsInt property is set to true.
type ID struct {
	ID    string
	IsInt bool
}

// UnmarshalJSON implements the json.Unmarshaler interface for the ID type. If
// the ID is an int, it is converted to a string and IsInt is set to true.
func (id *ID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("id cannot be empty")
	}

	if data[0] == '"' {
		var s string
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*id = ID{ID: s, IsInt: false}
	} else {
		var s int
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*id = ID{ID: strconv.Itoa(s), IsInt: true}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface for the ID type. If
// IsInt is set to true, then the ID is marshalled to an int rather than a
// string.
func (id ID) MarshalJSON() (out []byte, err error) {
	if id.IsInt {
		v, err := strconv.Atoi(id.ID)
		if err != nil {
			return nil, err
		}
		out, err = json.Marshal(v)
	} else {
		out, err = json.Marshal(id.ID)
	}

	return
}
