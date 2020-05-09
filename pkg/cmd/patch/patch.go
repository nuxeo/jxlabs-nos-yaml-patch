package patch

import (
	"fmt"
	"os"

	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/nxmatic/jxlabs-nos-helmfile-patch/pkg/yaml"
	"github.com/spf13/cobra"
)

type PatchOptions struct {
	*opts.CommonOptions
	Inplace string
}

func NewCmdPatch(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &PatchOptions{
		CommonOptions: commonOpts,
	}
	command := &cobra.Command{
		Use:     "patch",
		Short:   "patch json",
		Example: "patch patch.json",
		Args:    cobra.MinimumNArgs(1),
		Run: func(command *cobra.Command, args []string) {
			options.Cmd = command
			options.Args = args
			options.Run()
		},
	}

	command.Flags().StringVarP(&options.Inplace, "inplace", "i", "", "the json file which will be modified in place")

	return command
}

func (o *PatchOptions) Run() {
	if o.Inplace != "" {
		_, err := yaml.PatchInplace(o.Inplace, o.Args)
		helper.CheckErr(err)
	} else {
		bsTarget, err := yaml.Patch(os.Stdin, o.Args)
		helper.CheckErr(err)
		fmt.Printf("%s", string(bsTarget))
	}
}
