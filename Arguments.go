package main

import (
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	SrcPath string
	DstPath string
	Force   bool
	Verbose bool
}

func (s *Settings) ArgumentsVerif() error {

	flag.StringVar(&s.SrcPath, "s", "", "Dossier Photos")
	flag.StringVar(&s.DstPath, "d", "", "Dossier destination")
	flag.BoolVar(&s.Force, "f", false, "Force la creation Dossier destination")
	flag.BoolVar(&s.Verbose, "v", false, "Afficher details")
	flag.Parse()

	_, err := os.Stat(s.SrcPath)
	if err != nil {
		return fmt.Errorf("Dossier Photo Incorrect")
	}
	_, err = os.Stat(s.DstPath)
	if err != nil {
		rep := "Y"
		if !s.Force {
			fmt.Println("Dossier Destination Innexistant. Le créer ? Y or N")
			fmt.Scanln(&rep)
		}
		if rep == "Y" {
			err = os.Mkdir(s.DstPath, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
