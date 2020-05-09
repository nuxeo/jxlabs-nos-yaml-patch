package root

import (
	"fmt"
	"os"

	"github.com/nxmatic/jxlabs-nos-helmfile-patch/cmd/root/diff"
	"github.com/nxmatic/jxlabs-nos-helmfile-patch/cmd/root/patch"

	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/opts"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json-patch",
	Short: "compute diff and apply patch",
	Long:  ``,
}

// Execute executes the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	f := clients.NewFactory()
	commonOptions := opts.NewCommonOptionsWithTerm(f, os.Stdin, os.Stdout, os.Stderr)

	diffCmd := diff.NewCmdDiff(commonOptions)
	rootCmd.AddCommand(diffCmd)

	patchCmd := patch.NewCmdPatch(commonOptions)
	rootCmd.AddCommand(patchCmd)

	cobra.OnInitialize(func() {
		var err error
		if err != nil {
			panic(err)
		}
	})
}
