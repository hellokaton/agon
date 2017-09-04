package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func BasePath() string {
	file, _ := exec.LookPath(os.Args[0])
	fpath, _ := filepath.Abs(file)
	basePath := path.Dir(fpath)
	return basePath
}

func CreateFile(dir string, name string) (string, error) {
	src := path.Join(dir, name)

	_, err := os.Stat(src)

	if os.IsExist(err) {
		return src, nil
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		if os.IsPermission(err) {
			panic("You don't have permission to create files")
		}
		return "", err
	}

	_, err = os.Create(src)
	if err != nil {
		return "", err
	}

	return src, nil
}

func MkDir(filepath string) error {

	if _, err := os.Stat(filepath); err != nil {

		if os.IsNotExist(err) {

			err = os.MkdirAll(filepath, os.ModePerm)

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func CopyDir(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func WriteFile(file string, text string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Errorf("Open file error: %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.Write([]byte(text))
	if err != nil {
		return err
	}
	return w.Flush()
}
