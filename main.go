package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	CheminPhotos, Destination := arguments()
	VerifCheminPhotos, VerifDestination := verification(CheminPhotos, Destination)
	if VerifCheminPhotos || VerifDestination {
		return
	}
	//CheminPhotos := "./Photos"
	//Destination := "./Rangee"
	err := filepath.Walk(CheminPhotos, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			extention := strings.Split(path, ".")
			lenExtention := len(extention)
			if extention[lenExtention-1] != "mp4" {

				date := date_img(path)
				dateFormate := date.Format("2006-01-02")
				//fmt.Println(path)
				RepertoireDate(dateFormate, Destination, path)
			}
		}
		//fmt.Println(path)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
func RepertoireDate(date string, chemin string, photos string) {
	//fmt.Println("ici")
	TabDate := strings.Split(date, "-")
	TabPhoto := strings.Split(photos, "/")
	NbRepertoires := len(TabPhoto)
	//fmt.Println(TabDate)
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
	source, err := os.Open(photos)
	if err != nil {
		fmt.Println(err)
	}
	defer source.Close()
	dst, err := os.Create(CheminProvisoir + "/" + TabPhoto[NbRepertoires-1])
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()
	io.Copy(dst, source)
	//fmt.Println("fin")
}
