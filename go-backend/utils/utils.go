package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// MapStructure convert one type to another
func MapStructure(in, out interface{}) error {
	j, err := json.Marshal(in)
	if err != nil {
		return err
	}

	return json.Unmarshal(j, out)
}

// DirExists checks if a directory exists
func DirExists(p string) (bool, error) {
	if !filepath.IsAbs(p) {
		wd, err := os.Getwd()
		if err != nil {
			return false, err
		}
		p = filepath.Join(wd, p)
	}

	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
