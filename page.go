package main

import (
	humanize "github.com/dustin/go-humanize"
	"github.com/shopspring/decimal"
)

func loadPaymentTotals(roa Payments, payments Payments, paytotals *PaymentTotals) error {
	paytotals.PaidOut = payments.PaidOut.Total.Mul(decimal.NewFromFloat(-1))
	paytotals.NumPaidOut = payments.PaidOut.Num
	paytotals.NetCashDep = payments.Cash.Total.Sub(paytotals.PaidOut)
	paytotals.NumNetCash = payments.Cash.Num + payments.PaidOut.Num
	paytotals.CashCheckDep = paytotals.NetCashDep.Add(payments.Check.Total)
	paytotals.NumCashCheckDep = paytotals.NumNetCash + payments.Check.Num
	paytotals.CCDep = payments.CCPayments.Total.Add(payments.CCPulled.Total)
	paytotals.NumCCDep = payments.CCPayments.Num + payments.CCPulled.Num
	paytotals.TotalDailyDep = paytotals.CashCheckDep.Add(paytotals.CCDep)
	paytotals.NumTotalDailyDep = paytotals.NumCashCheckDep + paytotals.NumCCDep
	paytotals.TotalNonCash = payments.ARAdj.Total.Add(payments.OnAccount.Total.Add(payments.GiftCard.Total.Add(payments.Other.Total.Add(payments.Rewards.Total))))
	paytotals.NumTotalNonCash = payments.ARAdj.Num + payments.OnAccount.Num + payments.GiftCard.Num + payments.Other.Num + payments.Rewards.Num
	paytotals.GrandTotal = paytotals.TotalDailyDep.Add(paytotals.TotalNonCash)
	paytotals.ROAPay = roa.CCPayments.Total.Add(roa.CCPulled.Total.Add(roa.Cash.Total.Add(roa.Check.Total.Add(roa.GiftCard.Total.Add(roa.OnAccount.Total.Add(roa.Other.Total.Add(roa.PaidOut.Total.Add(roa.Rewards.Total))))))))
	paytotals.ROAARAdj = roa.ARAdj.Total

	temp := paytotals.GrandTotal.Add(paytotals.PaidOut)
	temp = temp.Sub(paytotals.ROAPay)
	temp = temp.Sub(paytotals.ROAARAdj)

	paytotals.PaymentsSales = temp

	return nil
}

func generatePage(transtot Totals, roa Payments, payments Payments, paytot PaymentTotals) error {
	var dispPay DispPayments
	generatePaymentDisplay(roa, payments, paytot, &dispPay)

	p.Totals = transtot
	p.Payments = dispPay
	p.ZeroMoney = "0.00"

	return nil
}

func generatePaymentDisplay(roa Payments, payments Payments, paytot PaymentTotals, dispPay *DispPayments) {
	dispPay.Cash = commaString(payments.Cash.Total, 2)
	dispPay.NumCash = humanize.Commaf(float64(payments.Cash.Num))
	dispPay.PaidOut = commaString(paytot.PaidOut, 2)
	dispPay.NumPaidOut = humanize.Commaf(float64(paytot.NumPaidOut))
	dispPay.NetCashDep = commaString(paytot.NetCashDep, 2)
	dispPay.Checks = commaString(payments.Check.Total, 2)
	dispPay.NumChecks = humanize.Commaf(float64(payments.Check.Num))
	dispPay.CashCheckDep = commaString(paytot.CashCheckDep, 2)
	dispPay.NumCashCheckDep = humanize.Commaf(float64(paytot.NumCashCheckDep))
	dispPay.CCDep = commaString(paytot.CCDep, 2)
	dispPay.NumCCDep = humanize.Commaf(float64(paytot.NumCCDep))
	dispPay.TotalDailyDep = commaString(paytot.TotalDailyDep, 2)
	dispPay.NumTotalDailyDep = humanize.Commaf(float64(paytot.NumTotalDailyDep))
	dispPay.OnAccount = commaString(payments.OnAccount.Total, 2)
	dispPay.NumOnAccount = humanize.Commaf(float64(payments.OnAccount.Num))
	dispPay.RewardsRedeem = commaString(payments.Rewards.Total, 2)
	dispPay.NumRewardsRedeem = humanize.Commaf(float64(payments.Rewards.Num))
	dispPay.ARAdj = commaString(payments.ARAdj.Total, 2)
	dispPay.NumARAdj = humanize.Commaf(float64(payments.ARAdj.Num))
	dispPay.GiftCards = commaString(payments.GiftCard.Total, 2)
	dispPay.NumGiftCards = humanize.Commaf(float64(payments.GiftCard.Num))
	dispPay.Other = commaString(payments.Other.Total, 2)
	dispPay.NumOther = humanize.Commaf(float64(payments.Other.Num))
	dispPay.TotalNonCash = commaString(paytot.TotalNonCash, 2)
	dispPay.NumTotalNonCash = humanize.Commaf(float64(paytot.NumTotalNonCash))
	dispPay.GrandTotal = commaString(paytot.GrandTotal, 2)
	dispPay.ROAOnAccount = commaString(paytot.ROAPay, 2)
	dispPay.ROAARAdj = commaString(paytot.ROAARAdj, 2)
	dispPay.PaymentsSales = commaString(paytot.PaymentsSales, 2)
}
