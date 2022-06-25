package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	s := Settings{}
	err := s.ArgumentsVerif()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s.SrcPath)

	deb := time.Now()

	if err != nil {
		fmt.Println(err)
	}

	err = s.run()
	fin := time.Now()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Photos Triées en :", fin.Sub(deb))
}

func (s *Settings) run() error {
	NbPhotos := 0
	err := filepath.Walk(s.SrcPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			extention := strings.ToLower(filepath.Ext(path))
			if extention != ".mp4" {

				date := date_img(path)
				dateFormate := date.Format("2006-01-02")
				RepertoireDate(dateFormate, s.DstPath, path)

				NbPhotos += 1
			}
		}
		//fmt.Println(path)

		return nil
	})
	return err
}

func RepertoireDate(date string, chemin string, photos string) {
	//fmt.Println("ici")
	TabDate := strings.Split(date, "-")
	TabDate = FormatageMois(TabDate)
	TabPhoto := strings.Split(photos, "/")
	NbRepertoires := len(TabPhoto)
	//fmt.Println(TabDate)
	CheminProvisoir := chemin
	_, err := os.Stat(CheminProvisoir)
	if err != nil {
		os.Mkdir(CheminProvisoir, os.ModePerm)
	}

	for i := 0; i < 3; i++ {
		CheminProvisoir += "/" + TabDate[i]
		_, err := os.Stat(CheminProvisoir)
		if err != nil {
			//fmt.Println(err)
			os.Mkdir(CheminProvisoir, os.ModePerm)
		}

	}
	source, _ := os.Open(photos)
	defer source.Close()
	dst, err := os.Create(CheminProvisoir + "/" + TabPhoto[NbRepertoires-1])
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()
	io.Copy(dst, source)
	//fmt.Println("fin")
}
