package virtual

import (
	"io/fs"
	"path"
	"time"
)

// A fileInfo implements fs.FileInfo and fs.DirEntry for a given map file.
type fileInfo struct {
	path    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	sys     interface{}
}

func (i *fileInfo) Name() string               { return path.Base(i.path) }
func (i *fileInfo) Mode() fs.FileMode          { return i.mode }
func (i *fileInfo) Type() fs.FileMode          { return i.mode.Type() }
func (i *fileInfo) ModTime() time.Time         { return i.modTime }
func (i *fileInfo) IsDir() bool                { return i.mode&fs.ModeDir != 0 }
func (i *fileInfo) Sys() interface{}           { return i.sys }
func (i *fileInfo) Info() (fs.FileInfo, error) { return i, nil }
func (i *fileInfo) Size() int64                { return i.size }
