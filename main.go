package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	CheminPhotos := "./Photos"
	CheminRange := "./Rangee"
	err := filepath.Walk(CheminPhotos, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			date := date_img(path)
			dateFormate := date.Format("2006-01-02")
			RepertoireDate(dateFormate, CheminRange, path)
		}
		//fmt.Println(path)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
func RepertoireDate(date string, chemin string, photos string) {
	TabDate := strings.Split(date, "-")
	fmt.Println(TabDate)
	CheminProvisoir := chemin
	_, err := os.Stat(CheminProvisoir)
	if err != nil {
		fmt.Println(err)
		os.Mkdir(CheminProvisoir, os.ModePerm)
	}

	for i := 0; i < 3; i++ {
		CheminProvisoir += "/" + TabDate[i]
		_, err := os.Stat(CheminProvisoir)
		if err != nil {
			fmt.Println(err)
			os.Mkdir(CheminProvisoir, os.ModePerm)
		}

	}

}

func date_img(fname string) time.Time {

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
	tm, err := x.DateTime()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(fname, tm)

	return tm
}
