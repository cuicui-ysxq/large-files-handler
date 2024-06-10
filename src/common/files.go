package common

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// If `singleFileSize` is 0, use `GitHubMaxRecommendedFileSize` as single file size;
// Otherwise, ensure that the file size is never larger than `GitHubMaxFileSize`.
func NormalizeSingleFileSize(singleFileSize FileSize) FileSize {
	if singleFileSize == 0 {
		return GitHubMaxRecommendedFileSize
	} else {
		return min(singleFileSize, GitHubMaxFileSize)
	}
}

func SplitFile(filePath string, chunkSize uint, outDir string) (outFilePaths []string, err error) {
	// ensure the single file size is never greater than GitHub's limits
	chunkSize = NormalizeSingleFileSize(chunkSize)

	// open input file
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// ensure that the input file is a regular file
	stat, err := file.Stat()
	if err != nil {
		return
	}
	if !stat.Mode().IsRegular() {
		err = fmt.Errorf("not a regular file: %s", filePath)
		return
	}

	// ensure that the output directory exists
	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		return
	}

	// split the input file
	basename := filepath.Base(filePath)
	buff := make([]byte, chunkSize)
	for i, eof := uint(0), false; !eof && (err == nil); i++ {
		// use an inner function so that
		// each output file is closed after writing to it has been completed
		err = func() (err error) {
			// read input file
			n, err := file.Read(buff)
			if err != nil {
				if err == io.EOF {
					eof = true
				} else {
					return
				}
			}
			if n == 0 {
				return
			}

			// create output file
			outFilePath := filepath.Join(outDir, fmt.Sprintf("%s.%s%d", basename, SplitSuffix, i))
			fmt.Println(outFilePath)
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
