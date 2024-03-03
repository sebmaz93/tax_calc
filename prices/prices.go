package prices

import (
	"fmt"
	"github.com/sebmaz93/tax_calc/conversion"
	"github.com/sebmaz93/tax_calc/iomanager"
	"math"
)

type PriceWithTaxJob struct {
	IO           iomanager.IO       `json:"-"`
	TaxRate      float64            `json:"tax_rate"`
	InputPrices  []float64          `json:"input_prices"`
	PriceWithTax map[string]float64 `json:"price_with_tax"`
}

func (job *PriceWithTaxJob) LoadData() error {
	lines, err := job.IO.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *PriceWithTaxJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.0f", price)] = math.Round((price*(1+job.TaxRate))*100) / 100
	}

	job.PriceWithTax = result
	return job.IO.WriteResult(job)
}

func NewPriceWithTaxJob(io iomanager.IO, taxRate float64) *PriceWithTaxJob {
	return &PriceWithTaxJob{
		IO:          io,
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
