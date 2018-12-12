package osutil

import (
	"io/ioutil"
	"path/filepath"
)

// CreateSimLink creates a simulated link which is just a text file containing the real path
// dst is the file name to point to
// src is the file name of the link
func CreateSimLink(dst, src string) error {
	dataBytes := []byte(dst)
	err := ioutil.WriteFile(src, dataBytes, 0644)
	return err
}

// ReadSimLink a simulated link which is just a text file containing the real path
// link is the file name of the link
func ReadSimLink(link string) (string, error) {
	path := ""
	bytes, err := ioutil.ReadFile(link)
	if err == nil {
		path = string(bytes)
		if !filepath.IsAbs(path) {
			path, err = filepath.Abs(filepath.Join(filepath.Dir(link), path))
		}
	}
	return path, err
}
