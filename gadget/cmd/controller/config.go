package main

import (
	"github.com/spf13/pflag"
	"os"
)

func init () {
	pflag.Parse()
}

var (
	kubeconfig = pflag.String("kubeconfig", os.Getenv("HOME") + "/.kube", "path custom kubeconfig")
)