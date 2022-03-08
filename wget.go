package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/cheggaaa/pb/v3"
)

func Wget(link string) func() error {
	return func() error {
		response, err := http.Get(link)
		if err != nil {
			return err
		}
		filename := GetFilename(link)
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		filesize := response.ContentLength
		bar := pb.New(int(filesize))
		bar.Start()
		reader := bar.NewProxyReader(response.Body)
		_, err = io.Copy(f, reader)
		bar.Finish()
		return err
	}
}

func GetFilename(link string) string {
	var fileName = filepath.Base(link)
	response, err := http.Get(link)
	if err != nil {
		return fileName
	}
	defer response.Body.Close()
	var p = regexp.MustCompile(`.+filename="(.+?)".*`)
	var contentDisposition = response.Header.Get(`Content-Disposition`)
	if contentDisposition != "" {
		matched := p.FindAllStringSubmatch(contentDisposition, -1)
		if len(matched) > 0 {
			return matched[0][0]
		}
	}
	return fileName
}
