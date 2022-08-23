package main

import (
	"bytes"
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
	err := s.VerifArguments()
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
	err := filepath.Walk(s.SrcPath, func(img string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		switch strings.ToUpper(filepath.Ext(img)) {
		case ".JPG", ".JPEG", ".CR2", ".CR3":
			return s.TraiteFichier(img)
		}
		return nil
	})
	return err
}

func (s *Settings) TraiteFichier(img string) error {
	file, err := os.Open(img)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	teeR := io.TeeReader(file, buffer)

	var date time.Time
	switch strings.ToUpper(filepath.Ext(img)) {
	case ".JPG", ".JPEG", ".CR2":
		date, err = DateJPG(teeR)
	case ".CR3":
		date, err = DateCR3(teeR)
	}

	if err != nil {
		return err
	}

	multiR := io.MultiReader(buffer, file)

	dir, err := s.DestinationDir(date)
	if err != nil {
		return err
	}

	destFileName := filepath.Join(dir, filepath.Base(img))
	w, err := os.Create(destFileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, multiR)

	return err
}

func (s *Settings) DestinationDir(date time.Time) (string, error) {
	datef := date.Format("2006/01/02")
	dest := filepath.Join(s.DstPath, datef)
	return dest, os.MkdirAll(dest, 0750)
}
