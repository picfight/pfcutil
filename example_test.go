package pfcutil_test

import (
	"fmt"
	"math"

	"github.com/picfight/pfcutil"
)

func ExampleAmount() {

	a := pfcutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = pfcutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = pfcutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 PFC
	// 100,000,000 Satoshis: 1 PFC
	// 100,000 Satoshis: 0.001 PFC
}

func ExampleNewAmount() {
	amountOne, err := pfcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := pfcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := pfcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := pfcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 PFC
	// 0.01234567 PFC
	// 0 PFC
	// invalid picfightcoin amount
}

func ExampleAmount_unitConversions() {
	amount := pfcutil.Amount(44433322211100)

	fmt.Println("Satoshi to kPFC:", amount.Format(pfcutil.AmountKiloPFC))
	fmt.Println("Satoshi to PFC:", amount)
	fmt.Println("Satoshi to MilliPFC:", amount.Format(pfcutil.AmountMilliPFC))
	fmt.Println("Satoshi to MicroPFC:", amount.Format(pfcutil.AmountMicroPFC))
	fmt.Println("Satoshi to Satoshi:", amount.Format(pfcutil.AmountSatoshi))

	// Output:
	// Satoshi to kPFC: 444.333222111 kPFC
	// Satoshi to PFC: 444333.222111 PFC
	// Satoshi to MilliPFC: 444333222.111 mPFC
	// Satoshi to MicroPFC: 444333222111 μPFC
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
