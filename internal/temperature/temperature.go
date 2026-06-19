package temperature

import "math"

func CelsiusToFahrenheit(c float64) float64 {
	f := c*1.8 + 32
	return math.Round(f*10) / 10
}

func CelsiusToKelvin(c float64) float64 {
	k := c + 273.15
	return math.Round(k*100) / 100
}
