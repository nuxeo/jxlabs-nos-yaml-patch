package diff

import (
	"io/ioutil"
	"log"

	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/mattbaird/jsonpatch"
	"github.com/spf13/cobra"
)

type PatchOptions struct {
	*opts.CommonOptions
}

func NewCmdPatch(commonOptions opts.CommonOptions) *cobra.Command {
	options := &PatchOptions{
		CommonOptions: commonOpts,
	}
	command := &cobra.Command{
		Use:     "patch",
		Short:   "patch json",
		Example: "patch patch.json",
		Args:    cobra.MinimumNArgs(2),
		Run: func(command *cobra.Command, args []string) {
			options.Cmd = command
			options.Args = args
			options.Run()
		},
	}

	command.Flags().StringVarP(&options.Original, "input", "i", "/dev/stdin", "the original json file")
	command.Flags().StringVarP(&options.Target, "output", "o", "/dev/stdout", "the target ouptput json file")

	return command
}

func (o *PatchOptions) Run() {
	patches := make([]jsonpatch.Patch, len(o.PatchFilePaths))

	for i, patchFilePath := range o.Args {
		var bs []byte
		bs, err = ioutil.ReadFile(patchFilePath)
		if err != nil {
			log.Fatalf("error reading patch file: %s", err)
		}

		patch, err = jsonpatch.DecodePatch(bs)
		if err != nil {
			log.Fatalf("error decoding patch file: %s", err)
		}

		patches[i] = patch
	}
}
