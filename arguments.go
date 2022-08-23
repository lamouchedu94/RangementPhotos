package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Settings struct {
	SrcPath string
	DstPath string
	Force   bool
	// Verbose bool
}

func (s *Settings) VerifArguments() error {

	flag.StringVar(&s.SrcPath, "s", "", "Dossier source")
	flag.StringVar(&s.DstPath, "d", "", "Dossier destination")
	flag.BoolVar(&s.Force, "f", false, "Force la creation Dossier destination")
	// flag.BoolVar(&s.Verbose, "v", false, "Afficher details")
	flag.Parse()

	_, err := os.Stat(s.SrcPath)
	if err != nil {
		return fmt.Errorf("Dossier source incorrect: %w", err)
	}
	_, err = os.Stat(s.DstPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Dossier destination incorrect: %w", err)
		}
		if !s.Force {
			rep := "Y"
		loop:
			for {
				fmt.Println("Dossier de destination innexistant. Le créer ? Y or N")
				fmt.Scanln(&rep)
				switch strings.ToUpper(rep) {
				case "Y":
					break loop
				case "N":
					os.Exit(1)
				}
			}
		}
		err = os.MkdirAll(s.DstPath, 0750)
		if err != nil {
			return err
		}
	}
	return nil
}
