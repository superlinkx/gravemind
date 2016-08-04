package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

func loadTransactions(transTotals *Totals) error {
	transactions := []Transaction{}
	var statInfo os.FileInfo
	var lastMod int64
	emptyFile := false
	misalignedFile := false

	f, err := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

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

		if err = gocsv.UnmarshalFile(f, &transactions); err != nil {
			fmt.Printf("Error unmarshalling file: %v\n", err)
			misalignedFile = true
		}

		if !misalignedFile {
			if err := checkTransDate(transactions[0].Date); err != nil {
				//sendErrorMail(err) //Disabled due to running every request. Will reimplement in Gravemind2
				p.TransError = fmt.Sprintf("Wrong date in file. Please report this error to your admin at %v", settings.ToEmail)
			}

			calcTotals(transactions, transTotals)

			p.TransLastMod = lastMod
			p.TransWarning = ""

			return nil
		}
	}

	emptyTotals(transTotals)

	p.TransWarning = "Empty set. Either no data yet or problem with transaction import. Contact Sysadmin if persists."
	p.TransLastMod = lastMod

	return errors.New("Transactions is an empty set")
}

func checkTransDate(transdate string) error {
	currentDate := time.Now()

	if transdate == currentDate.Format("01/02/06") {
		return nil
	}

	return errors.New("Date in transaction does not match current date.")
}
