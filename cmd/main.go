package main

import (
	"os"

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"

	"github.com/nxmatic/jxlabs-nos-yaml-patch/pkg/cmd"
)

func main() {
	f := clients.NewFactory()
	commonOptions := opts.NewCommonOptionsWithTerm(f, os.Stdin, os.Stdout, os.Stderr)
	cmd.NewCmd(commonOptions).Execute()
}
