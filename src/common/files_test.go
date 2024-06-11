package common_test

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"testing"

	"github.com/cuicui-yxsq/large-files-handler/common"
	"gotest.tools/assert"
)

const (
	numUint8Values uint16 = math.MaxUint8 + 1
	chunkSize      uint16 = 1 << 5

	randomFilePath string = "/dev/urandom"
)

func GetFileAndDir() (file, dir string) {
	_, file, _, _ = runtime.Caller(1)
	dir = filepath.Dir(file)

	return
}

func TestSplitZeroSizeChunk(t *testing.T) {
	_, err := common.SplitFile("", 0, "")
	assert.ErrorContains(t, err, "chunk size must be greater than 0")
}

func generateUint8Bytes() (allData [numUint8Values]byte, smallChunks [][chunkSize]byte) {
	var smallBuff [chunkSize]byte
	for i := uint16(0); i < numUint8Values; i++ {
		v := uint8(i)

		// use `allData` to store all bytes
		allData[i] = v

		// use `smallChunks` to store chunks of bytes
		pos := i % chunkSize
		if pos == 0 {
			smallBuff = [chunkSize]byte{}
		}
		smallBuff[pos] = v
		if pos == chunkSize-1 {
			smallChunks = append(smallChunks, smallBuff)
		}
	}
	return
}

func TestSplitFile(t *testing.T) {
	_, dir := GetFileAndDir()
	testDir := filepath.Join(dir, "../test")

	allData, smallChunks := generateUint8Bytes()

	// open and write the `allData` file
	allDataFilePath := filepath.Join(testDir, "all-data")
	{
		allDataFile, err := os.Create(allDataFilePath)
		assert.NilError(t, err)
		defer os.Remove(allDataFilePath)
		defer allDataFile.Close()

		allDataFile.Write(allData[:])
	}

	{ // split the `allData` file, and compare them with the `smallChunks`
		outFilePaths, err := common.SplitFile(allDataFilePath, uint(chunkSize), testDir)
		assert.NilError(t, err)

		assert.Equal(t, len(outFilePaths), len(smallChunks))
		for i, outFilePath := range outFilePaths {
			outBuff, err := os.ReadFile(outFilePath)
			defer os.Remove(outFilePath)

			assert.NilError(t, err)
			assert.Assert(t, slices.Equal(smallChunks[i][:], outBuff), "slices not equal")
		}
	}
}

func getRandomData(size uint, outFilePath string) (data []byte, err error) {
	data = make([]byte, size)

	randomFile, err := os.Open(randomFilePath)
	if err != nil {
		return
	}
	defer randomFile.Close()

	n, err := randomFile.Read(data)
	if err != nil {
		return
	}
	if uint(n) != size {
		err = fmt.Errorf("read %d bytes, expected %d bytes", n, size)
		return
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	n, err = outFile.Write(data)
	if err != nil {
		return
	}
	if uint(n) != size {
		err = fmt.Errorf("written %d bytes, expected %d bytes", n, size)
		return
	}

	return
}

func TestSplitRandomData(t *testing.T) {
	singleChunkSize := 100 * common.MB
	totalDataSize := singleChunkSize * 5

	testDir := ""

	// write random data to the test input file
	testInputFile := filepath.Join(testDir, "random")
	defer os.Remove(testInputFile)
	data, err := getRandomData(totalDataSize, testInputFile)
	assert.NilError(t, err)

	// get output files
	outFilePaths, err := common.SplitFile(testInputFile, singleChunkSize, testDir)
	defer func() {
		for _, outFilePath := range outFilePaths {
			os.Remove(outFilePath)
		}
	}()
	assert.NilError(t, err)
	// check the output files
	offset := uint(0)
	for _, outFilePath := range outFilePaths {
		buffer, err := os.ReadFile(outFilePath)
		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(data[offset:offset+singleChunkSize], buffer), "slices not equal")

		offset += singleChunkSize
	}
}

func TestMergeFiles(t *testing.T) {
	_, dir := GetFileAndDir()
	testDir := filepath.Join(dir, "../test")

	tests := []struct {
		inputFilename string
		content       []byte
	}{
		{"digits", []byte("0123456789")},
		{"capital-letters", []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")},
		{"small-letters", []byte("abcdefghijklmnopqrstuvwxyz")},
	}
	inputFilePaths := make([]string, len(tests))
	totalBytes := 0
	for i, test := range tests {
		// create input file
		inputFilePath := filepath.Join(testDir, test.inputFilename)
		inputFilePaths[i] = inputFilePath
		defer os.Remove(inputFilePath)

		err := os.WriteFile(inputFilePath, test.content, 0644)
		assert.NilError(t, err)
		totalBytes += len(test.content)
	}

	outputFilePath := filepath.Join(testDir, "merged")
	defer os.Remove(outputFilePath)

	// merge files
	n, err := common.MergeFiles(inputFilePaths, 128, outputFilePath)
	assert.NilError(t, err)
	assert.Equal(t, n, uint(totalBytes))
}
