package prices

import (
	"fmt"

	"cloudterms.net/price-cal/conversion"
	"cloudterms.net/price-cal/iomanager"
)

type taxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *taxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		withTaxPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", withTaxPrice)
	}

	fmt.Println(result)
	job.TaxIncludedPrices = result

	fmt.Println(job)
	job.IOManager.WriteJSON(job)
}

func (job *taxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadFileContents()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *taxIncludedPriceJob {
	return &taxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   io,
	}
}
