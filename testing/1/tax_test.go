package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected tax for %.2f to be %.2f, got %.2f", amount, expected, result)
	}
	amount = 1500.0
	expected = 10.0
	result = CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected tax for %.2f to be %.2f, got %.2f", amount, expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount   float64
		expected float64
	}

	cases := []calcTax{
		{amount: 500.0, expected: 5.0},
		{amount: 1500.0, expected: 10.0},
		{amount: 1000.0, expected: 10.0},
		{amount: 999.99, expected: 5.0},
	}

	for _, v := range cases {
		result := CalculateTax(v.amount)
		if result != v.expected {
			t.Errorf("Expected tax for %.2f to be %.2f, got %.2f", v.amount, v.expected, result)
		}
	}

}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	f.Add(500.0)
	f.Add(1500.0)
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount < 0 && result != 0.0 {
			t.Errorf("Expected tax for negative amount %.2f to be 0.0, got %.2f", amount, result)
		} else if amount >= 1000 && result != 10.0 {
			t.Errorf("Expected tax for %.2f to be 10.0, got %.2f", amount, result)
		} else if amount < 1000 && result != 5.0 {
			t.Errorf("Expected tax for %.2f to be 5.0, got %.2f", amount, result)
		}
	})
}
