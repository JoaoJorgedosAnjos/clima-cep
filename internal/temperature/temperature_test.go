package temperature

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	celsius := 28.5
	expectedFahrenheit := 83.3
	result := CelsiusToFahrenheit(celsius)

	if math.Abs(result-expectedFahrenheit) > 0.01 {
		t.Errorf("Erro na conversão para Fahrenheit. Esperado: %.1f, Obtido: %v", expectedFahrenheit, result)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	celsius := 28.5
	expectedKelvin := 301.65
	result := CelsiusToKelvin(celsius)

	if math.Abs(result-expectedKelvin) > 0.01 {
		t.Errorf("Erro na conversão para Kelvin. Esperado: %.2f, Obtido: %v", expectedKelvin, result)
	}
}
