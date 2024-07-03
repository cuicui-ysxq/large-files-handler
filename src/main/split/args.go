package main

import (
	"flag"
	"fmt"

	"github.com/cuicui-ysxq/large-files-handler/common"
)

type Args struct {
	inputFilePath string
	chunkSize     uint
	outDir        string
}

func (a *Args) DefineArgs() {
	flag.StringVar(&a.inputFilePath, "i", "", "input file path")
	flag.UintVar(&a.chunkSize, "s", common.GitHubMaxRecommendedFileSize, "chunk size in bytes")
	flag.StringVar(&a.outDir, "d", "", "output directory")
}

func (a *Args) Check() (errs []error) {
	if a.inputFilePath == "" {
		errs = append(errs, fmt.Errorf("missing input file path"))
	}

	return
}
