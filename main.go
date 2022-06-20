package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	chemin := "./Photos"

	err := filepath.Walk(chemin, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			date_img(path)
		}
		//fmt.Println(path)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
func date_img(fname string) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	/*
		camModel, _ := x.Get(exif.Model)
		fmt.Println(camModel.StringVal())
	*/
	tm, _ := x.DateTime()
	fmt.Println(fname, tm)
}
