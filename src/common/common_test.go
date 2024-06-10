package common_test

import (
	"path/filepath"
	"runtime"
)

func GetFileAndDir() (file, dir string) {
	_, file, _, _ = runtime.Caller(1)
	dir = filepath.Dir(file)

	return
}
