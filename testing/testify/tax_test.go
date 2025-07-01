package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	assert.Equal(t, expected, result, "Expected tax for %.2f to be %.2f, got %.2f", amount, expected, result)

	amount = 1500.0
	expected = 10.0
	result = CalculateTax(amount)
	assert.Equal(t, expected, result, "Expected tax for %.2f to be %.2f, got %.2f", amount, expected, result)
}

func TestCalculateTaxAndSave(t *testing.T) {
	mockRepo := new(TaxRepositoryMock)
	// 	mockRepo := &TaxRepositoryMock{}

	mockRepo.On("SaveTax", 5.0).Return(nil).Once()
	mockRepo.On("SaveTax", 10.0).Return(nil)
	mockRepo.On("SaveTax", 0.0).Return(errors.New("negative amount not allowed"))

	err := CalculateTaxAndSave(500.0, mockRepo)
	assert.NoError(t, err)

	err = CalculateTaxAndSave(1500.0, mockRepo)
	assert.NoError(t, err)

	err = CalculateTaxAndSave(-100.0, mockRepo)
	assert.Error(t, err, "Expected error for negative amount")

	// Verifique todas as expectativas apenas no final
	mockRepo.AssertExpectations(t)
}
