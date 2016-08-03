package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

func loadPayments(payments *Payments) error {
	var paymentData []Payment
	var statInfo os.FileInfo
	var lastMod int64
	emptyFile := false
	misalignedFile := false

	f, err := os.OpenFile(settings.PaymentsFile, os.O_RDONLY, os.ModePerm)

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

		if err = gocsv.UnmarshalFile(f, &paymentData); err != nil {
			fmt.Printf("Error unmarshalling file: %v\n", err)
			misalignedFile = true
		}

		if !misalignedFile {
			parsePayment(paymentData, payments)

			p.PaymentsLastMod = lastMod
			p.PaymentsWarning = ""

			return nil
		}
	}

	emptyPayment(payments)
	p.PaymentsWarning = "Empty set. Either no data yet or problem with Payments import. Contact Sysadmin if persists."
	p.PaymentsLastMod = lastMod

	return errors.New("Payments is an empty set")
}
