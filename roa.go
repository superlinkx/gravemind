package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

func loadROA(p *Page) error {
	var roa Payments
	var statInfo os.FileInfo
	var lastMod int64
	emptyFile := false
	misalignedFile := false

	f, err := os.OpenFile(settings.ROAFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		emptyFile = true
	}

	defer f.Close()

	if !emptyFile {
		var e error
		statInfo, e = f.Stat()
		if e != nil {
			fmt.Printf("Error getting statinfo: %v\n", err)
			lastMod = 0
		} else {
			lastMod = statInfo.ModTime().Unix()
			currentTime := time.Now().Unix()

			lastMod = (currentTime - lastMod) / 60
		}

		if err = gocsv.UnmarshalFile(f, &roa); err != nil {
			fmt.Printf("Error unmarshalling file: %v\n", err)
			misalignedFile = true
		}

		if !misalignedFile {
			calcROA(roa)

			p.ROALastMod = lastMod
			p.ROAWarning = ""

			return nil
		}
	}

	emptyROA()
	p.ROAWarning = "Empty set. Either no data yet or problem with ROA import. Contact Sysadmin if persists."
	p.ROALastMod = lastMod

	return errors.New("ROA is an empty set")
}
