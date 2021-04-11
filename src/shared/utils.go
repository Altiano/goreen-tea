package shared

import (
	"fmt"
)

func MakeHTTPError(code int, msg error) error {
	return fmt.Errorf("%v: %w", code, msg)
}

func TabArrow() {
	fmt.Print("-> ")
}
