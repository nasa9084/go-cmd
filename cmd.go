package cmd

import (
	"io/ioutil"
	"strings"

	"golang.org/x/sys/unix"
)

// DfStat is a data type of df.
type DfStat struct {
	All  uint64
	Free uint64
	Used uint64
}

// Df is `df`
func Df(path string) (DfStat, error) {
	if path == "" {
		path = "/"
	}
	fs := unix.Statfs_t{}
	if err := unix.Statfs(path, &fs); err != nil {
		return DfStat{}, err
	}
	all := fs.Blocks * uint64(fs.Bsize) / 1024
	free := fs.Bfree * uint64(fs.Bsize) / 1024
	return DfStat{
		All:  all,
		Free: free,
		Used: all - free,
	}, nil
}

func ls(path string, all bool) ([]string, error) {
	if path == "" {
		path = "."
	}
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, fi := range fis {
		if all || !strings.HasPrefix(fi.Name(), ".") {
			names = append(names, fi.Name())
		}
	}
	return names, nil
}

// Ls is `ls`
func Ls(path string) ([]string, error) {
	return ls(path, false)
}

// LsAll is `ls -A`
func LsAll(path string) ([]string, error) {
	return ls(path, true)
}
