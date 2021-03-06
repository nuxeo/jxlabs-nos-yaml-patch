package cmd

import (
	"github.com/jenkins-x/jx/v2/pkg/cmd/opts"
	"github.com/jenkins-x/jx-logging/pkg/log"

	"github.com/nuxeo/jxlabs-nos-yaml-patch/pkg/cmd/apply"
	"github.com/nuxeo/jxlabs-nos-yaml-patch/pkg/cmd/generate"

	"github.com/spf13/cobra"
)

func NewCmd(commonOpts *opts.CommonOptions) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "patch-yaml",
		Short: "[command] patch-yaml",
		Long:  `compute [diff] and apply [patch]`,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}

	cmd.AddCommand(generate.NewCmd(commonOpts))
	cmd.AddCommand(apply.NewCmd(commonOpts))

	cobra.OnInitialize(func() {
		var err error
		if err != nil {
			panic(err)
		}
	})

	return cmd
}
