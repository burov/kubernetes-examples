package unknown

import (
	"errors"
)

type Command struct{}

func (Command) Execute(args []string) error {
	return errors.New("unable to find command")
}
