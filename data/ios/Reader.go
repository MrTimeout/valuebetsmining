package ios

import (
	"errors"
	"os"
	"path/filepath"
)

//Restart ... Remove file if it exists and creates a new one
func Restart(name string) error {
	src, err := os.Stat(name)
	if os.IsNotExist(err) {
		errDir := os.Mkdir(name, 0775)
		if errDir != nil {
			return errDir
		}
	} else if !os.IsNotExist(err) {
		err := os.Remove(name)
		if err != nil {
			return err
		}
		err = os.Mkdir(name, 0775)
		if err != nil {
			return err
		}
	}
	if src.Mode().IsRegular() {
		return errors.New("Already exists like a file")
	}
	return nil
}

//RemoveContents ... Removes all of a dir
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
