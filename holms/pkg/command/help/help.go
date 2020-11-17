package help

import (
	"errors"
	"fmt"
)

type Command struct {}

var template = `
	holms help you to find pods assigned to some resources, can be used as a Kubernetes plugin

	Commands:
		pods - list all pods assigned to particular service, deployment, daemonset, replicaset, statefulset
			args: 
				* --namespace - namespace for resource
				* --resource  - resource selector, have to be specified in format {type}/{name}
			example:
				kubectl-holms pods --namespace custom-ns --resource=deployments/app
`

func (Command) Execute(args []string) error {
	fmt.Print(template)
	return errors.New("unimplemented")
}
