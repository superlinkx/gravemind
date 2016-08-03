package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	humanize "github.com/dustin/go-humanize"
	"github.com/shopspring/decimal"
)

//Globals
var settings Server
var p *Page

func handler(w http.ResponseWriter, r *http.Request) {
	var transTotals Totals
	var roa Payments
	var payments Payments
	var paytotals PaymentTotals

	p.BusinessName = settings.BusinessName

	if err := loadTransactions(&transTotals); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := loadROA(&roa); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := loadPayments(&payments); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := loadPaymentTotals(roa, payments, &paytotals); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := generatePage(transTotals, roa, payments, paytotals); err != nil {
		fmt.Printf("%v\n", err)
	}

	t, err := template.ParseFiles(settings.DashboardTemplate)
	if err != nil {
		error500(&w, "Error parsing dashboard template", err)
		return
	}

	t.Execute(w, p)
}

func commaString(dec decimal.Decimal, places int32) string {
	floatConv, _ := strconv.ParseFloat(dec.StringFixed(places), 64)
	return humanize.Commaf(floatConv)
}

func error500(w *http.ResponseWriter, msg string, err error) {
	http.Error(*w, msg, 500)
	fmt.Printf("%v: %v\n", msg, err)
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
