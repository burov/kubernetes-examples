package main

import (
	"fmt"
	"os"

	"github.com/burov/kubernetes-examples/holms/pkg/command"
)

func main() {
	fmt.Println(os.Args)
	cmd := command.GetCommand(os.Args[1])
	if err := cmd.Execute(os.Args[2:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
