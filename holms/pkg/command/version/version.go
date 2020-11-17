package version

import "fmt"

const version = "v0.0.1"

type Command struct {}

func (c Command) Execute(args []string) error {
	fmt.Println(version)
	return nil
}
