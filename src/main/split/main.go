package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cuicui-yxsq/large-files-handler/common"
)

var args struct {
	inputFilePath string
	chunkSize     uint
	outDir        string
}

func init() {
	flag.StringVar(&args.inputFilePath, "i", "", "input file path")
	flag.UintVar(&args.chunkSize, "s", 0, "chunk size")
	flag.StringVar(&args.outDir, "d", "", "output directory")

	flag.Parse()

	errs := []error{}
	if args.inputFilePath == "" {
		errs = append(errs, fmt.Errorf("missing input file path"))
	}

	checkAndPrintErrsThenExit(errs, 1)
}

func checkAndPrintErrsThenExit(errs []error, exitCode int) {
	if len(errs) > 0 {
		flag.Usage()

		fmt.Fprintln(os.Stderr, "Error(s):")
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "-", err)
		}
		os.Exit(exitCode)
	}
}

func main() {
	outFiles, err := common.SplitFile(args.inputFilePath, args.chunkSize, args.outDir)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Output file(s): %d file(s)", len(outFiles))
	for _, outFile := range outFiles {
		fmt.Fprintln(os.Stderr, "-", outFile)
	}
}
