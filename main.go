package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {

	//VerifCheminPhotos, VerifDestination := verification(CheminPhotos, Destination)
	//CheminPhotos := "./Photos"
	//Destination := "./Rangee"
	err := ArgumentsVerif()
	if err != nil {
		fmt.Println(err)
		return
	}
	deb := time.Now()
	var TabPathPhoto []string
	var TabDate []string
	var wg sync.WaitGroup
	settings := Settings{}
	err = filepath.Walk(settings.SrcPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			extention := strings.Split(path, ".")
			lenExtention := len(extention)
			if extention[lenExtention-1] != "mp4" {

				date := DateImg(path)
				dateFormate := date.Format("2006-01-02")
				TabDate = append(TabDate, dateFormate)
				TabPathPhoto = append(TabPathPhoto, path)

			}
		}
		return err
	})
	Debut := 0
	NbPhotos := 0
	interval := decoupage(TabPathPhoto)
	echantillon := interval

	for i := 0; i < 4; i++ {
		if echantillon > len(TabPathPhoto) {
			echantillon = Debut + (len(TabPathPhoto) - Debut)
		}
		wg.Add(1)
		go func(Destination string, Debut int, echantillon int) {
			Boucle(Debut, echantillon, TabDate, Destination, TabPathPhoto)
			wg.Done()
		}(settings.DstPath, Debut, echantillon)
		//fmt.Println("Deb", Debut, "Echan", echantillon, "inter", interval)
		Debut += interval
		echantillon += interval

	}
	wg.Wait()

	if err != nil {
		fmt.Println(err)
	}
	fin := time.Now()
	//fmt.Println(len(TabPathPhoto), len(TabDate))
	fmt.Println(NbPhotos, "Photos Triées en :", fin.Sub(deb))
}

func decoupage(TabPathPhoto []string) int {
	long := len(TabPathPhoto)
	return (long / 4) + 1 //runtime.NumCPU()) + 1
}

func Boucle(debut int, fin int, TabDate []string, Destination string, path []string) {
	fmt.Println(debut, fin, "ici")
	for i := debut; i < fin; i++ {
		RepertoireDate(TabDate[i], Destination, path[i])
	}
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
