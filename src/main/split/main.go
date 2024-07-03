package main

import (
	"fmt"
	"os"

	"github.com/cuicui-yxsq/large-files-handler/args"
	"github.com/cuicui-yxsq/large-files-handler/common"
)

func main() {
	var a Args
	args.ParseAndCheckArgs(&a, 1)

	outFiles, err := common.SplitFile(a.inputFilePath, a.chunkSize, a.outDir)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "Output file(s): %d file(s)", len(outFiles))
	for _, outFile := range outFiles {
		fmt.Fprintln(os.Stderr, "-", outFile)
	}
}
