package cmd

import (
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/log"

	"github.com/nxmatic/jxlabs-nos-helmfile-patch/pkg/cmd/diff"
	"github.com/nxmatic/jxlabs-nos-helmfile-patch/pkg/cmd/patch"

	"github.com/spf13/cobra"
)

func NewCmd(commonOpts *opts.CommonOptions) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "patch-helmfile",
		Short: "compute diff and apply patch",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}

	diffCmd := diff.NewCmdDiff(commonOpts)
	cmd.AddCommand(diffCmd)

	patchCmd := patch.NewCmdPatch(commonOpts)
	cmd.AddCommand(patchCmd)

	cobra.OnInitialize(func() {
		var err error
		if err != nil {
			panic(err)
		}
	})

	return cmd
}
