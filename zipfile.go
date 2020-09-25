/*
* Utility functions for handling and fetching repo archives in zip format
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:26
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 18:12:04
 */
package quick

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// UnZip Download and unpack a zipfile at a given URI.
//
// This will download the zipfile to the cookiecutter repository,
// and unpack into a temporary directory.
//
// :param zipURI: The URI for the zipfile.
// :param isURL: Is the zip URI a URL or a file?
// :param cloneToDir: The cookiecutter repository directory
//                    to put the archive into.
// :param noInput: Suppress any prompts
// :param password: The password to use when unpacking the repository.

func UnZipFromURL(zipURI string, isURL bool, cloneToDir string, noInput bool, password string)error{

}

func UnZip(cloneToDir string, noInput bool, password string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
