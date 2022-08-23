package main

import (
	"bytes"
	"io"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func DateJPG(r io.Reader) (time.Time, error) {
	ex, err := exif.Decode(r)

	if err != nil {
		return time.Time{}, err
	}
	return ex.DateTime()
}

func DateCR3(r io.Reader) (time.Time, error) {
	buff := make([]byte, 0x400)
	_, err := io.ReadFull(r, buff)
	if err != nil {
		return time.Time{}, err
	}

	i := bytes.Index(buff, []byte("CMT1"))
	if i < 0 {
		return time.Time{}, nil
	}

	ex, err := exif.Decode(bytes.NewReader(buff[i+4:]))
	if err != nil {
		return time.Time{}, err
	}
	return ex.DateTime()
}
