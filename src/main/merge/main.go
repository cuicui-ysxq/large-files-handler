package main

import (
	"fmt"
	"os"

	"github.com/cuicui-yxsq/large-files-handler/args"
	"github.com/cuicui-yxsq/large-files-handler/common"
)

func main() {
	var a Args
	if hasErrs := args.ParseAndCheckArgs(&a); hasErrs {
		os.Exit(1)
	}

	n, err := common.MergeFiles(a.inputFilePaths, a.bufferSize, a.outputFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Merged %d file(s) into output file \"%s\", total: %d byte(s)\n", len(a.inputFilePaths), a.outputFilePath, n)
}
