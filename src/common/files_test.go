package common_test

import (
	"math"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/cuicui-yxsq/large-files-handler/common"
	"gotest.tools/assert"
)

const (
	numUint8Values uint16 = math.MaxUint8 + 1
	chunkSize      uint16 = 1 << 5
)

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
