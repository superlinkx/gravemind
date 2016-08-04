package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func emptyTotals(transtot *Totals) {
	emptyTime, _ := time.Parse("3:04 pm", "12:00 am")
	emptyTimeString := emptyTime.Format("3:04 pm")
	zeromoney := "0.00"
	zero := "0"

	transtot.Sales = zeromoney
	transtot.Tax = zeromoney
	transtot.Total = zeromoney
	transtot.InvCount = zero
	transtot.InvPerHr = zeromoney
	transtot.SalesPerHr = zeromoney
	transtot.SalesPerInv = zeromoney
	transtot.FirstTransTime = emptyTimeString
	transtot.LastTransTime = emptyTimeString
	transtot.Cost = zeromoney
	transtot.Profit = zeromoney
}

func emptyPayment(payment *Payments) {
	var zeroMoney decimal.Decimal
	zeroMoney = decimal.NewFromFloat(0)
	const zero = 0
	payment.PaidOut.Total = zeroMoney
	payment.CCPayments.Total = zeroMoney
	payment.Cash.Total = zeroMoney
	payment.Check.Total = zeroMoney
	payment.Rewards.Total = zeroMoney
	payment.Other.Total = zeroMoney
	payment.OnAccount.Total = zeroMoney
	payment.GiftCard.Total = zeroMoney
	payment.ARAdj.Total = zeroMoney
	payment.PaidOut.Num = zero
	payment.CCPayments.Num = zero
	payment.Cash.Num = zero
	payment.Check.Num = zero
	payment.Rewards.Num = zero
	payment.Other.Num = zero
	payment.OnAccount.Num = zero
	payment.GiftCard.Num = zero
	payment.ARAdj.Num = zero
}

func calcTotals(transactions []Transaction, transtot *Totals) error {
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

	perprofit := profit.Div(salesTotal).Mul(decimal.NewFromFloat(100))

	transtot.Sales = commaString(salesTotal, 2)
	transtot.Tax = commaString(taxTotal, 2)
	transtot.Total = commaString(grandTotal, 2)
	transtot.InvCount = commaString(invCount, 0)
	transtot.InvPerHr = commaString(invPerHr, 4)
	transtot.SalesPerHr = commaString(salesPerHr, 2)
	transtot.SalesPerInv = commaString(salesPerInv, 2)
	transtot.FirstTransTime = leastTime.Format("3:04 pm")
	transtot.LastTransTime = greatestTime.Format("3:04 pm")
	transtot.Cost = commaString(costTotal, 2)
	transtot.Profit = commaString(profit, 2)
	transtot.PerProfit = perprofit.StringFixed(1)

	return nil
}

func parsePayment(paymentdata []Payment, payment *Payments) {
	for _, pay := range paymentdata {
		switch pay.PaymentName {
		case "Paid Out (Cash)":
			payment.PaidOut.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.PaidOut.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Credit Card Payments":
			payment.CCPayments.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.CCPayments.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Credit Card Pulled Early":
			payment.CCPulled.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.CCPulled.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Cash":
			payment.Cash.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.Cash.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Check":
			payment.Check.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.Check.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Cellar Rats Rewards":
			payment.Rewards.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.Rewards.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "On Account":
			payment.OnAccount.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.OnAccount.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "Gift Card":
			payment.GiftCard.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.GiftCard.Num, _ = strconv.Atoi(pay.TotalPayments)
		case "AR Adjustments":
			payment.ARAdj.Total, _ = decimal.NewFromString(pay.TotalAmount)
			payment.ARAdj.Num, _ = strconv.Atoi(pay.TotalPayments)
		default:
			newAmount, _ := decimal.NewFromString(pay.TotalAmount)
			payment.Other.Total = payment.Other.Total.Add(newAmount)
			newCount, _ := strconv.Atoi(pay.TotalPayments)
			payment.Other.Num += newCount
		}
	}
}
