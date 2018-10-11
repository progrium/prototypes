//go:generate go run -tags=gen generate.go
package assets

import (
	"io"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var RootPath string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	RootPath, _ = filepath.Abs(path.Join(path.Dir(filename), ".."))
}

func FindTemplate(skip int) (io.Reader, error) {
	_, filename, _, _ := runtime.Caller(skip)
	callerPath := strings.Replace(filename, RootPath+"/", "", -1)
	return Assets.Open(strings.Replace(callerPath, ".go", ".html", -1))
}
