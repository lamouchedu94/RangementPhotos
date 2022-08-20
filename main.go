package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	copyf "github.com/lamouchedu94/RangementPhotos/copy"
	"github.com/lamouchedu94/RangementPhotos/shoot"
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
	err := filepath.Walk(s.SrcPath, func(img string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		date, err := shoot.Date(img)
		if err != nil {
			return nil
		}

		dir, err := s.finaldir(date)
		if err != nil {
			return err
		}
		copyf.CopyPictures(img, dir)

		return nil
	})
	return err
}

func (s *Settings) finaldir(date time.Time) (string, error) {
	datef := date.Format("2006-01-02")
	tabdate := strings.Split(datef, "-")
	datef = s.DstPath
	for _, val := range tabdate {
		datef += "/" + val

		_, err := os.Stat(datef)

		if err != nil {
			err = os.Mkdir(datef, 0750)
			if err != nil {
				return "", err
			}
		}

	}
	return datef, nil
}
