package tax

import "time"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {

	tax := CalculateTax(amount)
	return repository.SaveTax(tax)

}

func CalculateTax(amount float64) float64 {

	if amount < 0 {
		return 0.0 // No tax for negative amounts
	}

	if amount >= 1000 {
		return 10.0 // 15% tax for amounts >= 1000
	}
	return 5.0
}

func CalculateTax2(amount float64) float64 {

	time.Sleep(time.Millisecond) // Simulate a delay for testing purposes

	if amount < 0 {
		return 0.0 // No tax for negative amounts
	}

	if amount >= 1000 {
		return 10.0 // 15% tax for amounts >= 1000
	}
	return 5.0
}
