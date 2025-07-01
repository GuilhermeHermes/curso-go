package tax

import "time"

func CalculateTax(amount float64) float64 {
	// guilhermehermes@pop-os:~/Documentos/curso-go/testing/1$ go test -coverprofile=coverage.out
	// PASS
	// coverage: 80.0% of statements
	// ok      github/GuilhermeHermes/learning_go/testing      0.004s
	// guilhermehermes@pop-os:~/Documentos/curso-go/testing/1$ go tool cover -html=coverage.out
	// guilhermehermes@pop-os:~/Documentos/curso-go/testing/1$
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

// guilhermehermes@pop-os:~/Documentos/curso-go/testing/1$ go test -bench=. -run=^#
// goos: linux
// goarch: amd64
// pkg: github/GuilhermeHermes/learning_go/testing
// cpu: AMD Ryzen 7 5700U with Radeon Graphics
// BenchmarkCalculateTax-16        1000000000               0.2378 ns/op
// BenchmarkCalculateTax2-16           1092           1098820 ns/op
// PASS
// ok      github/GuilhermeHermes/learning_go/testing      1.586s
