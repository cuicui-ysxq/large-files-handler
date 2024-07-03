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

func ParseAndCheckArgs(args Args, exitCode int) {
	args.DefineArgs()

	if !flag.Parsed() {
		flag.Parse()
	}

	errs := args.Check()
	if len(errs) > 0 {
		flag.Usage()

		fmt.Fprintln(os.Stderr, "Error(s):")
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "-", err)
		}
		os.Exit(exitCode)
	}
}
