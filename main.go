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

	start := time.Now()

	err = s.run()
	end := time.Now()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Photos Triées en :", end.Sub(start))
}

func (s *Settings) run() error {
	PicturesNb := 0
	err := filepath.Walk(s.SrcPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		extention := strings.ToLower(filepath.Ext(path))
		if extention != ".mp4" {

			date := date_img(path)
			dateFormate := date.Format("2006-01-02")
			CurrentPath, err := RepertoireDate(dateFormate, s.DstPath, path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			CopyPictures(path, CurrentPath)

			PicturesNb += 1
		}

		return nil
	})
	return err
}

func RepertoireDate(date string, chemin string, photos string) (string, error) {
	//fmt.Println("ici")

	TabDate := strings.Split(date, "-")
	TabDate = FormatageMois(TabDate)

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
	return CheminProvisoir, err
}
func CopyPictures(Pictures string, CurrentPath string) {
	TabPhoto := strings.Split(Pictures, "/")
	NbRepertoires := len(TabPhoto)

	source, _ := os.Open(Pictures)

	defer source.Close()
	dst, err := os.Create(CurrentPath + "/" + TabPhoto[NbRepertoires-1])
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()

	io.Copy(dst, source)
	dst.Close()
	//fmt.Println("fin")
}
