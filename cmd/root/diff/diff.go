package diff

import (
	"fmt"
	"io/ioutil"

	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/mattbaird/jsonpatch"
	"github.com/spf13/cobra"
)

type DiffOptions struct {
	*opts.CommonOptions
}

func NewCmdDiff(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &DiffOptions{
		CommonOptions: commonOpts,
	}
	command := &cobra.Command{
		Use:     "diff",
		Short:   "generate json patch from the diff between two json objects",
		Example: "diff file1.json file2.json",
		Args:    cobra.MinimumNArgs(2),
		Run: func(command *cobra.Command, args []string) {
			options.Cmd = command
			options.Args = args
			options.Run()
		},
	}

	return command
}

func (o *DiffOptions) Run() {
	bsOriginal, err := ioutil.ReadFile(o.Args[0])
	helper.CheckErr(err)
	bsTarget, err := ioutil.ReadFile(o.Args[1])
	helper.CheckErr(err)
	patch, err := jsonpatch.CreatePatch(bsOriginal, bsTarget)
	helper.CheckErr(err)
	for _, operation := range patch {
		fmt.Printf("%s\n", operation.Json())
	}
}
