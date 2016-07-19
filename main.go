package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	sr "github.com/superlinkx/gravemind/simpleround"

	"github.com/gocarina/gocsv"
)

//Global for settings
var settings Server

// Page Structure
type Page struct {
	BusinessName string
	LastMod      int64
	Totals       Totals
}

// Totals storage
type Totals struct {
	Sales          string
	Tax            string
	Total          string
	InvCount       string
	InvPerHr       string
	SalesPerHr     string
	SalesPerInv    string
	FirstTransTime string
	LastTransTime  string
	Cost           string
	Profit         string
}

// Transaction Structure for CSVData
type Transaction struct {
	TransID  string `csv:"Trans#"`
	Date     string `csv:"Date"`
	Time     string `csv:"Time"`
	SubTotal string `csv:"SubTot"`
	Tax      string `csv:"Tax"`
	Total    string `csv:"Total"`
	Cost     string `csv:"Cost"`
	NumTrans string `csv:"#Trns"`
}

// Server Parameters Structure
type Server struct {
	BusinessName      string `json:"businessname"`
	TransactionsFile  string `json:"transfile"`
	DashboardTemplate string `json:"dashboard_template"`
	Port              string `json:"port"`
}

// Arguments Structure
type Arguments struct {
	Config string
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{}

	f, err := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer f.Close()

	statInfo, err := f.Stat()
	if err != nil {
		fmt.Printf("Error getting statinfo: %v\n", err)
		return
	}
	lastMod := statInfo.ModTime().Unix()
	currentTime := time.Now().Unix()

	lastMod = (currentTime - lastMod) / 60

	if err = gocsv.UnmarshalFile(f, &transactions); err != nil {
		fmt.Printf("Error unmarshalling file: %v\n", err)
		return
	}

	totals := calcTotals(transactions)

	p := &Page{
		BusinessName: settings.BusinessName,
		LastMod:      lastMod,
		Totals:       totals,
	}

	t, err := template.ParseFiles(settings.DashboardTemplate)
	if err != nil {
		fmt.Printf("Error parsing template file: %v\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, p)
}

func calcTotals(transactions []Transaction) Totals {
	var invCount float64
	var salesTotal float64
	var taxTotal float64
	var grandTotal float64
	var costTotal float64

	var sales float64
	var tax float64
	var total float64
	var cost float64

	greatestTime, _ := time.Parse("3:04 pm", "12:00 am")
	leastTime, _ := time.Parse("3:04 pm", "11:59 pm")

	invCount = float64(len(transactions))

	for _, transaction := range transactions {
		sales, _ = strconv.ParseFloat(transaction.SubTotal, 64)
		tax, _ = strconv.ParseFloat(transaction.Tax, 64)
		total, _ = strconv.ParseFloat(transaction.Total, 64)
		cost, _ = strconv.ParseFloat(transaction.Cost, 64)

		salesTotal += sales
		taxTotal += tax
		grandTotal += total
		costTotal += cost

		transTime, _ := time.Parse("3:04 pm", strings.TrimSpace(transaction.Time))

		if transTime.After(greatestTime) {
			greatestTime = transTime
		}

		if transTime.Before(leastTime) {
			leastTime = transTime
		}
	}

	totalTime := greatestTime.Sub(leastTime).Hours()

	invPerHr := invCount / totalTime
	salesPerHr := salesTotal / totalTime
	salesPerInv := salesTotal / invCount

	profit := salesTotal - costTotal

	return Totals{
		Sales:          humanize.Commaf(sr.RoundDollars(salesTotal)),
		Tax:            humanize.Commaf(sr.RoundDollars(taxTotal)),
		Total:          humanize.Commaf(sr.RoundDollars(grandTotal)),
		InvCount:       humanize.Commaf(invCount),
		InvPerHr:       humanize.Commaf(sr.RoundDollars(invPerHr)),
		SalesPerHr:     humanize.Commaf(sr.RoundDollars(salesPerHr)),
		SalesPerInv:    humanize.Commaf(sr.RoundDollars(salesPerInv)),
		FirstTransTime: leastTime.Format("3:04 pm"),
		LastTransTime:  greatestTime.Format("3:04 pm"),
		Cost:           humanize.Commaf(sr.RoundDollars(costTotal)),
		Profit:         humanize.Commaf(sr.RoundDollars(profit)),
	}
}

func getArgs(params *Arguments) error {
	flag.StringVar(&params.Config, "config", "/etc/gravemind/gravemind.json", "Location of config file (default is /etc/gravemind/gravemind.json)")

	flag.Parse()

	return nil
}

func readConfig(config string, server *Server) error {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return err
	}

	if err := json.Unmarshal(file, &server); err != nil {
		fmt.Printf("JSON Unmarshalling error: %v\n", err)
		return err
	}

	return nil
}

func main() {
	var args Arguments
	if err := getArgs(&args); err != nil {
		fmt.Printf("Error getting arguments: %v\n", err)
	}

	if err := readConfig(args.Config, &settings); err != nil {
		fmt.Printf("Error reading configuration: %v\n", err)
	}

	address := ":" + settings.Port

	fmt.Println(address)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
