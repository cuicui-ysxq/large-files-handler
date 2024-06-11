package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cuicui-yxsq/large-files-handler/common"
)

type inputFilePaths []string

func (i *inputFilePaths) Set(s string) error {
	*i = append(*i, s)
	return nil
}

func (i *inputFilePaths) String() string {
	return fmt.Sprint(*i)
}

var args struct {
	inputFilePaths inputFilePaths
	outputFilePath string

	bufferSize uint
}

func init() {
	flag.Var(&args.inputFilePaths, "i", "input files")
	flag.StringVar(&args.outputFilePath, "o", "", "output file path")

	flag.UintVar(&args.bufferSize, "b", common.DefaultReadBufferSize, "buffer size")

	flag.Parse()

	errs := []error{}
	if len(args.inputFilePaths) == 0 {
		errs = append(errs, fmt.Errorf("no input files specified"))
	}
	if args.outputFilePath == "" {
		errs = append(errs, fmt.Errorf("missing output file path"))
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
	n, err := common.MergeFiles(args.inputFilePaths, args.bufferSize, args.outputFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Merged %d file(s) into output file \"%s\", total: %d byte(s)\n", len(args.inputFilePaths), args.outputFilePath, n)
}
