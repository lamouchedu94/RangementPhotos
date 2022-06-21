package main

import (
	"flag"
	"fmt"
	"os"
)

func arguments() (string, string) {
	ArgCheminPhotos := flag.String("c", "", "Dossier Photos")
	ArgDestination := flag.String("d", "", "Dossier destination")
	flag.Parse()
	return *ArgCheminPhotos, *ArgDestination
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
