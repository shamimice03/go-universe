package main

import (
	"fmt"

	"cloudterms.net/price-cal/cmdmanager"
	"cloudterms.net/price-cal/filemanager"
	"cloudterms.net/price-cal/prices"
)

func main() {
	inputFilePath := "prices.txt"
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	cmd := cmdmanager.New()

	// calculation from file
	for i, taxRate := range taxRates {
		outputFilePath := fmt.Sprintf("result_%0.f.json", taxRates[i]*100)
		fm := filemanager.New(inputFilePath, outputFilePath)

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

	// calculation from cmd
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Process()
	}
}
