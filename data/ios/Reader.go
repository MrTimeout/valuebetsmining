package ios

import (
	"errors"
	"os"
)

//Restart ... Remove file if it exists and creates a new one
func Restart(name string) error {
	src, err := os.Stat(name)
	if os.IsNotExist(err) {
		errDir := os.Mkdir(name, 0775)
		if errDir != nil {
			return errDir
		}
	}
	if src.Mode().IsRegular() {
		return errors.New("Already exists like a file")
	}
	return nil
}
