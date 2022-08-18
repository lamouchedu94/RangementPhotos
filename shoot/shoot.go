package shoot

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func Date(filename string) (time.Time, error) {
	r, err := os.Open(filename)
	if err != nil {
		return time.Time{}, nil
	}
	defer r.Close()
	d, err := exifdate(r)

	if err == nil {
		return d, nil
	}
	//fmt.Println(err)
	_, err = r.ReadAt(nil, 0) //A voir pour .CR3
	if err != nil {
		return time.Time{}, err
	}
	d, err = readCR3(filename)

	if err == nil {
		return d, err
	}
	return time.Time{}, nil
}

func exifdate(r io.Reader) (time.Time, error) {
	ex, err := exif.Decode(r)

	if err != nil {
		return time.Time{}, err
	}
	return ex.DateTime()
}

func readCR3(img string) (time.Time, error) {
	b, err := os.ReadFile(img)
	if err != nil {
		return time.Time{}, err
	}
	i := bytes.Index(b, []byte("CMT1"))
	//fmt.Println(i)
	r := bytes.NewBuffer(b[i+4:])
	d, err := decode(r)
	if err != nil {
		return time.Time{}, err
	}
	return d, err
}

func decode(r *bytes.Buffer) (time.Time, error) {

	ex, err := exif.Decode(r)
	if err != nil {
		return time.Time{}, err
	}
	d, err := ex.DateTime()
	return d, err
	/*
		if err != nil {
			return nil, err
		}
		for _, t := range ex.Tiff.Dirs {
			return t, err
		}
		return nil, nil
	*/
}
