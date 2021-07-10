package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
	function to unzip the zipped file
	pass the source and destination as parameters to call the function
	unzips the file to the destination path
*/
func Unzip(source, destination string) error {
	//use zip package to open and read the zip file
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer func() {
		// close the read file at the end of the function
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(destination, 0755)

	//Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWrite := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			err := rc.Close()
			if err != nil {
				panic(err)
			}
		}()

		//create the file path by joining the destination directory and file name
		path := filepath.Join(destination, f.Name)

		//check for directory traversal
		if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
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
				err := f.Close()
				if err != nil {
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

	//calls the extract and write function for all the files in the zipped folder
	for _, f := range r.File {
		err := extractAndWrite(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

/*
	function to check whether the zip file exists
	checking on the file info on the specified path
	if there are no errors, then file exists
*/
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
