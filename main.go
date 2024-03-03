package main

import (
	"fmt"
	"github.com/sebmaz93/tax_calc/cmdmanager"
	"github.com/sebmaz93/tax_calc/filemanager"
	"github.com/sebmaz93/tax_calc/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.20}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate*100))
		_ = cmdmanager.New()
		priceJob := prices.NewPriceWithTaxJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			return
		}
	}

}
