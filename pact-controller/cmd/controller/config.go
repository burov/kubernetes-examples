package main

import (
	"os"

	"github.com/spf13/pflag"
)

func init() {
	pflag.Parse()
}

var (
	kubeconfig = pflag.String("kubeconfig", os.Getenv("HOME") + "/.kube/config", "Path to kubeconfig file")
)
