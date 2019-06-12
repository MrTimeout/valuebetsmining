package ios

import (
	"log"
	"os"
	"path/filepath"
)

//Restart ... Remove dir if it exists and creates a new one
func Restart(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.MkdirAll(name, 0755)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = os.RemoveAll(name)
		if err != nil {
			log.Println(err)
		}
		err = os.MkdirAll(name, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

//RestartFile ... Remove file if it exists and creates a new one
func RestartFile(name string) error {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		err := os.Remove(name)
		if err != nil {
			return err
		}
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
