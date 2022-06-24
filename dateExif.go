package main

import (
	"log"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func DateImg(fname string) time.Time {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	/*
		camModel, _ := x.Get(exif.Model)
		fmt.Println(camModel.StringVal())
	*/
	tm, err := x.DateTime() //	fmt.Println(fname, tm)

	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(fname, tm)

	return tm
}
