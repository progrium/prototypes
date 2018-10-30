// +build gen

package assets

import (
	"net/http"
	"regexp"

	"github.com/spf13/afero"
)

var Assets http.FileSystem

func init() {
	fs := afero.NewRegexpFs(afero.NewBasePathFs(afero.NewOsFs(), RootPath), regexp.MustCompile(`\.html$`))
	httpFs := afero.NewHttpFs(fs)
	Assets = httpFs.Dir(".")
}
