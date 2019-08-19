package main

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func copyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func computeFileSHA1(file string) string {
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}

	hash := sha1.New()
	hash.Write(content)
	return hex.EncodeToString(hash.Sum(nil))
}

func computeContentSHA1(b []byte) string {
	hash := sha1.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func downloadToFile(url string, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("status != 200")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.Remove(file)
	if err != nil {
		return err
	}

	dest, err := os.Create(file)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = dest.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath,
				string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
