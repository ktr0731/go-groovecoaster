package groovecoaster

import (
	"fmt"
	"strconv"
)

// IntToBool type is a type that regards int as bool
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
