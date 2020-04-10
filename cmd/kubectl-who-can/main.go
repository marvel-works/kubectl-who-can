package main

import (
	"fmt"
	"github.com/aquasecurity/kubectl-who-can/pkg/cmd"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
	// Load all known auth plugins
	"flag"
	"github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog"
	"os"
)

func initFlags() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Hide all klog flags except for -v
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		if f.Name != "v" {
			pflag.Lookup(f.Name).Hidden = true
		}
	})
}

func main() {
	defer klog.Flush()

	initFlags()
	root, err := cmd.NewWhoCanCommand(clioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
