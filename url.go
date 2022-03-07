package main

import (
	"path/filepath"
	"strings"
)

func Join(prefix string, suffix string) string {
	if strings.HasSuffix(prefix, "/") {
		return prefix + suffix
	}
	return prefix + "/" + suffix
}

// RepoDir get directory name with git repo
// example:
//    1. https://github.com/hellojukay/httpfs.git  => httpfs
//    2. https://github.com/hellojukay/httpfs      => httpfs

func RepoDir(repo string) string {
	dir := filepath.Base(repo)
	if strings.HasSuffix(dir, ".git") {
		return strings.TrimSuffix(dir, ".git")
	}
	return dir
}
