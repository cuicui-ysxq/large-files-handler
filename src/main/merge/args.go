package main

import (
	"flag"
	"fmt"

	"github.com/cuicui-yxsq/large-files-handler/common"
)

type (
	Args struct {
		inputFilePaths inputFilePaths
		outputFilePath string

		bufferSize uint
	}

	inputFilePaths []string
)

func (a *Args) DefineArgs() {
	flag.Var(&a.inputFilePaths, "i", "input files")
	flag.StringVar(&a.outputFilePath, "o", "", "output file path")

	flag.UintVar(&a.bufferSize, "b", common.DefaultReadBufferSize, "buffer size")
}

func (a *Args) Check() (errs []error) {
	if len(a.inputFilePaths) == 0 {
		errs = append(errs, fmt.Errorf("no input files specified"))
	}
	if a.outputFilePath == "" {
		errs = append(errs, fmt.Errorf("missing output file path"))
	}

	return
}

func (i *inputFilePaths) Set(s string) error {
	*i = append(*i, s)
	return nil
}

func (i *inputFilePaths) String() string {
	return fmt.Sprint(*i)
}
