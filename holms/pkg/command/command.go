package command

import (
	"github.com/burov/kubernetes-examples/holms/pkg/command/help"
	"github.com/burov/kubernetes-examples/holms/pkg/command/pods"
	"github.com/burov/kubernetes-examples/holms/pkg/command/unknown"
	"github.com/burov/kubernetes-examples/holms/pkg/command/version"
)

type Command interface {
	Execute(args []string) error
}

var registery = map[string]Command{
	"version": version.Command{},
	"help":    help.Command{},
	"pods":    pods.Command{},
}

func GetCommand(name string) Command {
	cmd, ok := registery[name]
	if !ok {
		return unknown.Command{}
	}

	return cmd
}
