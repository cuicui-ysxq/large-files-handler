package common

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	DefaultReadBufferSize uint = 16 * MB
)

// Split files into chunks of the specified size.
func SplitFile(filePath string, chunkSize uint, outDir string) (outFilePaths []string, err error) {
	if chunkSize == 0 {
		err = fmt.Errorf("chunk size must be greater than 0")
		return
	}

	// open input file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer inputFile.Close()

	// ensure that the input file is a regular file
	stat, err := inputFile.Stat()
	if err != nil {
		return
	}
	if !stat.Mode().IsRegular() {
		err = fmt.Errorf("not a regular file: %s", filePath)
		return
	}

	// if `outDir` is not specified, use current working directory
	if outDir == "" {
		outDir = "."
	} else {
		// otherwise, ensure that the output directory exists
		// if it does not exist, create it
		err = os.MkdirAll(outDir, 0755)
		if err != nil {
			return
		}
	}

	// split the input file
	basename := filepath.Base(filePath)
	buff := make([]byte, chunkSize)
	for i, eof := uint(0), false; !eof && (err == nil); i++ {
		// use an inner function so that
		// each output file is closed after writing to it has been completed
		err = func() (err error) {
			// read input file
			n, err := inputFile.Read(buff)
			if err != nil {
				if err == io.EOF {
					eof = true
				} else {
					return
				}
			}
			if n == 0 {
				return // do not write empty files
			}

			// create output file
			outFilePath := filepath.Join(outDir, fmt.Sprintf("%s.%s%d", basename, SplitSuffix, i))
			outFile, err := os.Create(outFilePath)
			if err != nil {
				return
			}
			defer outFile.Close()

			// write output file
			_, err = outFile.Write(buff[:n])
			if err != nil {
				return
			}

			outFilePaths = append(outFilePaths, outFilePath)
			return
		}()
	}

	if err == io.EOF {
		err = nil
	}
	return
}

// Merge files into a single file.
// The contents of the input files are read into a buffer chunk by chunk, and then copied to the output file.
// If buffer size is set to 0, then the `DefaultReadBufferSize` is used.
func MergeFiles(filenames []string, bufferSize uint, outFilePath string) (n uint, err error) {
	if bufferSize == 0 {
		bufferSize = DefaultReadBufferSize
	}

	// create output file
	outputFile, err := os.Create(outFilePath)
	if err != nil {
		return
	}
	defer outputFile.Close()

	// read from each input file
	n = 0
	buff := make([]byte, bufferSize)
	for _, filePath := range filenames {
		// use an inner function so that
		// each output file is closed after reading from it has been completed
		err = func() (err error) {
			// open input file
			inputFile, err := os.Open(filePath)
			if err != nil {
				return
			}
			defer inputFile.Close()

			// ensure that the input file is a regular file
			stat, err := inputFile.Stat()
			if err != nil {
				return
			}
			if !stat.Mode().IsRegular() {
				err = fmt.Errorf("not a regular file: %s", filePath)
				return
			}

			// copy input file content to output file
			bytesCopied, err := io.CopyBuffer(outputFile, inputFile, buff)
			n += uint(bytesCopied)

			return
		}()
		if err != nil {
			return
		}
	}
	return
}
