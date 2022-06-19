package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	chemin := "/home"

	err := filepath.Walk(chemin, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
