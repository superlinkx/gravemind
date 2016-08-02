package main

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func emptyTotals() Totals {
	emptyTime, _ := time.Parse("3:04 pm", "12:00 am")
	emptyTimeString := emptyTime.Format("3:04 pm")
	return Totals{
		Sales:          "0.00",
		Tax:            "0.00",
		Total:          "0.00",
		InvCount:       "0",
		InvPerHr:       "0.000",
		SalesPerHr:     "0.00",
		SalesPerInv:    "0.00",
		FirstTransTime: emptyTimeString,
		LastTransTime:  emptyTimeString,
		Cost:           "0.00",
		Profit:         "0.00",
	}
}

func calcTotals(transactions []Transaction) Totals {
	var invCount decimal.Decimal
	var salesTotal decimal.Decimal
	var taxTotal decimal.Decimal
	var grandTotal decimal.Decimal
	var costTotal decimal.Decimal

	var sales decimal.Decimal
	var subtot decimal.Decimal
	var discount decimal.Decimal
	var ship decimal.Decimal
	var tax decimal.Decimal
	var total decimal.Decimal
	var cost decimal.Decimal

	greatestTime, _ := time.Parse("3:04 pm", "12:00 am")
	leastTime, _ := time.Parse("3:04 pm", "11:59 pm")

	invCount = decimal.NewFromFloat(float64(len(transactions)))

	for _, transaction := range transactions {
		subtot, _ = decimal.NewFromString(transaction.SubTotal)
		discount, _ = decimal.NewFromString(transaction.Discount)
		ship, _ = decimal.NewFromString(transaction.Ship)
		sales = subtot.Sub(discount).Add(ship)
		tax, _ = decimal.NewFromString(transaction.Tax)
		total, _ = decimal.NewFromString(transaction.Total)
		cost, _ = decimal.NewFromString(transaction.Cost)

		salesTotal = salesTotal.Add(sales)
		taxTotal = taxTotal.Add(tax)
		grandTotal = grandTotal.Add(total)
		costTotal = costTotal.Add(cost)

		transTime, _ := time.Parse("3:04 pm", strings.TrimSpace(transaction.Time))

		if transTime.After(greatestTime) {
			greatestTime = transTime
		}

		if transTime.Before(leastTime) {
			leastTime = transTime
		}
	}

	totalTime := greatestTime.Sub(leastTime).Hours()

	invPerHr := invCount.Div(decimal.NewFromFloat(totalTime))
	salesPerHr := salesTotal.Div(decimal.NewFromFloat(totalTime))
	salesPerInv := salesTotal.Div(invCount)

	profit := salesTotal.Sub(costTotal)

	return Totals{
		Sales:          commaString(salesTotal, 2),
		Tax:            commaString(taxTotal, 2),
		Total:          commaString(grandTotal, 2),
		InvCount:       commaString(invCount, 0),
		InvPerHr:       commaString(invPerHr, 4),
		SalesPerHr:     commaString(salesPerHr, 2),
		SalesPerInv:    commaString(salesPerInv, 2),
		FirstTransTime: leastTime.Format("3:04 pm"),
		LastTransTime:  greatestTime.Format("3:04 pm"),
		Cost:           commaString(costTotal, 2),
		Profit:         commaString(profit, 2),
	}
}
