package main

import (
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	SrcPath string
	DstPath string
}

func ArgumentsVerif() error {
	settings := Settings{}
	flag.StringVar(&settings.SrcPath, "s", "", "Dossier Photos")
	flag.StringVar(&settings.DstPath, "d", "", "Dossier destination")

	fmt.Println(settings.DstPath, settings.SrcPath)

	flag.Parse()
	_, err := os.Stat(settings.SrcPath)
	if err != nil {
		return fmt.Errorf("Dossier Photo Incorrect")
	}
	_, err = os.Stat(settings.DstPath)
	if err != nil {
		fmt.Println("Dossier Destination Innexistant. Le créer ? Y or N")
		rep := ""
		fmt.Scanln(&rep)
		if rep == "Y" {
			err = os.Mkdir(settings.DstPath, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func verification(CheminPhotos, destination string) (bool, bool) {
	photo := false
	dest := false
	_, err1 := os.Stat(CheminPhotos)
	if err1 != nil {
		fmt.Println("Dossier Photo Incorrect")
		photo = true
		return photo, true
	}
	_, err2 := os.Stat(destination)
	if err2 != nil {
		fmt.Println("Dossier Destination Innexistant. Le créer ? Y or N")
		rep := ""
		fmt.Scanln(&rep)
		if rep == "Y" {
			dest = false
		} else {
			dest = true
		}

	}
	return photo, dest
}
