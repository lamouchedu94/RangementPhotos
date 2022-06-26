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

func (s *Settings) ArgumentsVerif() error {

	flag.StringVar(&s.SrcPath, "s", "", "Dossier Photos")
	flag.StringVar(&s.DstPath, "d", "", "Dossier destination")
	flag.Parse()

	_, err := os.Stat(s.SrcPath)
	if err != nil {
		return fmt.Errorf("Dossier Photo Incorrect")
	}
	_, err = os.Stat(s.DstPath)
	if err != nil {
		fmt.Println("Dossier Destination Innexistant. Le créer ? Y or N")
		rep := ""
		fmt.Scanln(&rep)
		if rep == "Y" {
			err = os.Mkdir(s.DstPath, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
