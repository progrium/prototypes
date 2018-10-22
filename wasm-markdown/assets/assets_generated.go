// Code generated by vfsgen; DO NOT EDIT.

// +build !gen

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 10, 13, 18, 4, 51, 853655144, time.UTC),
		},
		"/app": &vfsgen۰DirInfo{
			name:    "app",
			modTime: time.Date(2018, 10, 11, 17, 13, 53, 729530237, time.UTC),
		},
		"/app/footer.html": &vfsgen۰FileInfo{
			name:    "footer.html",
			modTime: time.Date(2018, 10, 11, 18, 49, 55, 248536825, time.UTC),
			content: []byte("\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x66\x6f\x6f\x74\x65\x72\x22\x3e\x7b\x7b\x20\x43\x6f\x70\x79\x72\x69\x67\x68\x74\x20\x7d\x7d\x20\x3c\x73\x6c\x6f\x74\x3e\x3c\x2f\x73\x6c\x6f\x74\x3e\x3c\x2f\x64\x69\x76\x3e"),
		},
		"/app/pageview.html": &vfsgen۰CompressedFileInfo{
			name:             "pageview.html",
			modTime:          time.Date(2018, 10, 11, 18, 50, 12, 292425553, time.UTC),
			uncompressedSize: 354,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x51\x4b\xc3\x30\x14\x85\xdf\xf7\x2b\x2e\x79\x2f\xd9\x44\x50\x6a\x12\x90\x81\x38\x41\x7c\xf1\x0f\x64\xcb\xed\x16\xcc\x72\x43\x7a\x6d\x57\xc6\xfe\xbb\x2c\x5d\x45\x9b\x97\x84\x73\x4f\x38\xdf\x3d\x6a\x4b\x6e\x30\xb0\x00\x00\x50\xce\x77\xd0\xf2\x10\x50\x8b\x26\x90\xe5\x1a\xb2\xdf\x1f\x58\x98\x32\x2e\x16\xc6\x13\xdb\x8c\x16\x7e\xa5\xeb\xe9\x2a\x8a\xb5\x8f\xe9\x9b\xb5\xf8\x88\x9f\x78\xe2\xe7\x8c\x76\x7d\xb0\x71\x8f\xe2\xbf\x35\x53\xdf\x6a\xb1\xba\x9f\xc9\x3b\x0a\xad\x16\x0f\xcb\x99\x3c\xd1\x50\xe4\xaa\xb1\x47\x1f\x86\x1a\x8e\x14\xa9\x4d\x76\x87\x4f\xc2\x9c\xcf\x9b\x6b\xea\xe5\xa2\xe4\x44\x36\xc2\x2a\xe9\x7c\x77\x7b\xbe\xdb\xfc\xe5\xa8\x8f\xd0\x55\x5b\x1f\x5d\xbd\x19\x41\xcb\x25\x8c\x92\xd3\xfc\x66\x7f\x21\x62\xcc\xb0\xa6\x34\x94\xf5\xb5\xb8\x5b\xae\x1e\xe1\x0d\x9b\xe6\x6f\x13\xc9\xbc\x62\x08\x04\x3d\xe5\xe0\x94\x4c\x53\xee\xf8\xdd\xc0\x42\xc9\xd2\xed\x4f\x00\x00\x00\xff\xff\x42\x8a\x6d\x72\x62\x01\x00\x00"),
		},
		"/assets": &vfsgen۰DirInfo{
			name:    "assets",
			modTime: time.Date(2018, 10, 11, 17, 55, 4, 427936947, time.UTC),
		},
		"/dev": &vfsgen۰DirInfo{
			name:    "dev",
			modTime: time.Date(2018, 10, 11, 3, 59, 34, 848057779, time.UTC),
		},
		"/static": &vfsgen۰DirInfo{
			name:    "static",
			modTime: time.Date(2018, 10, 13, 18, 6, 3, 390808834, time.UTC),
		},
		"/static/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2018, 10, 11, 18, 11, 37, 827661288, time.UTC),
			uncompressedSize: 1152,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xd1\x6b\xdb\x3e\x10\xc7\x9f\xed\xbf\xe2\xea\x27\x07\x1a\xb9\xbf\xdf\x53\x68\x6c\x43\xdb\x95\x32\x18\x1b\x2c\x1b\x7d\x56\xe4\xb3\xad\x4e\xd6\x19\xe9\xdc\x60\x4a\xff\xf7\xa1\xa8\x49\x97\xc0\xd6\x07\xdb\x58\x77\xdf\xef\x9d\x3e\xd2\x95\x17\x0d\x29\x9e\x47\x84\x9e\x07\x53\xa7\xe5\xc5\x72\x99\xde\xd1\x38\x3b\xdd\xf5\x0c\xff\x5f\xfd\xb7\x82\x1f\x3d\xc2\x03\xc1\xcd\xc4\x3d\x39\x2f\xe0\xc6\x18\xd8\x87\x3d\x38\xf4\xe8\x9e\xb1\x11\xe9\x4f\x8f\x40\x2d\x70\xaf\x3d\x78\x9a\x9c\x42\x50\xd4\x20\x68\x0f\x1d\x3d\xa3\xb3\xd8\xc0\x76\x06\x09\xb7\x9b\x4f\x4b\xcf\xb3\xc1\xd4\x68\x85\xd6\x23\x70\x2f\x19\x94\xb4\xb0\x45\x68\x69\xb2\x0d\x68\x0b\xdc\x23\x7c\xf9\x7c\x77\xff\x75\x73\x0f\xad\x36\x28\xd2\xe5\xb2\x4e\xcb\xd8\x65\x5a\xf6\x28\x9b\x3a\x4d\xca\x01\x59\x82\xea\xa5\xf3\xc8\x55\x36\x71\xbb\x5c\x65\x61\x9d\x35\x1b\xac\x1f\x08\x76\xd2\x0f\x65\x11\x7f\xd3\x34\x29\xbd\x72\x7a\x64\xf0\x4e\x55\x59\x88\x89\x27\x9f\xd5\x65\x11\x97\xeb\x63\x42\x9d\x26\x89\x6e\x21\xbf\x78\xc4\xed\x8d\xf7\x38\x6c\xcd\x2c\xb4\xf5\x2c\x2d\x6b\xc9\xb8\x61\x87\x72\xd0\xb6\x5b\xc0\x0b\x14\x05\x8c\x64\xe6\x56\x1b\x93\x26\x49\xf2\x91\x04\x2a\x90\x7e\xb6\x0a\x72\x87\x7e\xbc\x04\x3d\x8c\xe4\xf8\xdb\xf6\x09\x15\x2f\xa0\xaa\xe1\x25\xb8\x24\x8a\xac\xe7\x03\xcb\x0a\xe4\x4e\x6a\x86\x3c\x7e\x82\x70\x21\xa4\x73\x72\xbe\x9d\xda\x16\x5d\xbe\x58\xef\x45\x0e\x79\x72\xf6\x2d\xf9\x2f\x8d\xe4\xd1\xf3\xac\xf0\x5e\xff\x1a\xde\xaf\x69\x7a\xa8\xde\x11\x54\x60\x71\x07\x0f\x14\x2b\x18\x64\x18\xa8\xb9\x84\xe0\xb7\x0e\x89\xed\x64\x15\x6b\xb2\xe0\xd0\x90\x6c\xf2\x45\x6c\xff\x23\x08\x79\x8b\xac\xfa\x3c\x93\xe3\x28\xc2\x39\x64\x8b\x4b\xe8\x48\x9c\xb4\x24\xb8\x47\x9b\xbf\xb3\x9a\xcc\x9f\x7c\x06\x6a\xa0\x82\xb8\x2c\x06\x6a\x26\x83\x91\x41\x28\xf7\x1e\x89\xc5\xd5\x5b\xac\x28\xc2\xc6\xc8\xa0\x50\x06\xe5\x11\x5b\xe4\xd5\x91\x70\x93\xcd\x83\x62\x71\x62\xf5\x6f\x9c\x7b\x1e\xe7\xbd\xaf\xc3\xad\x08\xe3\xc1\x70\xe8\x60\x0f\x38\x1a\x3f\x4b\x07\x3b\xff\xc6\xf6\x11\xb7\x1b\x52\xbf\x90\xf3\x6c\xe7\xaf\x8b\xc2\x90\x92\xa6\x27\xcf\xd7\xab\xab\xd5\x55\x11\xb9\x66\x51\xb8\xf3\x82\x2c\x8d\x68\xa1\x82\xfc\xec\xb2\x84\x5d\x19\xea\xf2\xec\xfb\x5e\x81\x0e\x1c\xca\x66\x16\x42\x04\xf1\xbe\xfa\xd1\x62\x40\xef\x65\x87\x67\x2e\x3b\x2f\x94\x21\x8f\xef\xb7\x29\x1e\xe9\xfa\x20\x0e\x4f\x98\xa2\xe3\xbc\x94\x45\x9c\xc4\xb4\xdc\x52\x33\x9f\xce\xcf\x41\x7d\x9a\x1f\xf3\x82\x30\x8c\xf2\xef\x00\x00\x00\xff\xff\xc8\x5c\xee\x1f\x80\x04\x00\x00"),
		},
		"/wasm": &vfsgen۰DirInfo{
			name:    "wasm",
			modTime: time.Date(2018, 10, 12, 20, 16, 59, 462076991, time.UTC),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app"].(os.FileInfo),
		fs["/assets"].(os.FileInfo),
		fs["/dev"].(os.FileInfo),
		fs["/static"].(os.FileInfo),
		fs["/wasm"].(os.FileInfo),
	}
	fs["/app"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app/footer.html"].(os.FileInfo),
		fs["/app/pageview.html"].(os.FileInfo),
	}
	fs["/static"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}