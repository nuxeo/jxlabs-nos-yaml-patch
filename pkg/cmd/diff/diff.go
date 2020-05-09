package diff

import (
	"fmt"

	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/nxmatic/jxlabs-nos-helmfile-patch/pkg/yaml"
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
	bsPatch, err := yaml.CreatePatchFromFiles(o.Args[0], o.Args[1])
	helper.CheckErr(err)
	fmt.Printf("%s\n", string(bsPatch))
}
