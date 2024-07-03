package args

import (
	"flag"
	"fmt"
	"os"
)

type Args interface {
	DefineArgs()
	Check() (errs []error)
}

func ParseAndCheckArgs(args Args) (hasErrs bool) {
	args.DefineArgs()

	if !flag.Parsed() {
		flag.Parse()
	}

	errs := args.Check()
	hasErrs = len(errs) > 0
	if hasErrs {
		flag.Usage()

		fmt.Fprintln(os.Stderr, "Error(s):")
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "-", err)
		}
	}

	return
}
