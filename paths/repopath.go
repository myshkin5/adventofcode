package paths

import (
	"path/filepath"
	"runtime"
)

func SourcePath() string {
	_, f, _, _ := runtime.Caller(1)

	return filepath.Dir(f)
}
