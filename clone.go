package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Clone(repo string) func() error {
	return func() error {
		c := exec.Command("git", "clone", repo)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	}
}

func SetOrigin(repo string) func() error {
	return func() error {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		dir := RepoDir(repo)
		dir = filepath.Join(pwd, dir)
		c := exec.Command("git", "remote", "set-url", "origin", repo)
		c.Dir = dir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	}
}
