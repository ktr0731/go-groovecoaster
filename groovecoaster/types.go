package groovecoaster

import (
	"fmt"
	"strconv"
)

// IntToBool type is a type that regards int as bool for JSON
type IntToBool bool

// UnmarshalJSON unmarshals int to bool
func (i *IntToBool) UnmarshalJSON(bytes []byte) error {
	n, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}

	switch n {
	case 1:
		*i = true
	case 0:
		*i = false
	default:
		return fmt.Errorf("Invalid value in IntToBool")
	}

	return nil
}

// StringToBool type is a type that regards string as bool for JSON
type StringToBool bool

// UnmarshalJSON unmarshals string to bool
func (s *StringToBool) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "1":
		*s = true
	case "0":
		*s = false
	default:
		return fmt.Errorf("Invalid value in StringToBool")
	}

	return nil
}
