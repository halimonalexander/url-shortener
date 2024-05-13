package e

import (
	"errors"
	"fmt"
)

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func IfIsChangeTo(err error, assert error, changeTo error) error {
	if errors.Is(err, assert) {
		return changeTo
	}
	return err
}
