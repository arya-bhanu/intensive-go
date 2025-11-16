package main

import "testing"

func TestCelciusToFahrenheit(t *testing.T) {

}

func TestFahrenheitToCelcius(t *testing.T) {
	tests := []struct {
		name            string
		fahrenheit      float64
		expectedCelcius float64
	}{
		{
			name:            "negative value test",
			fahrenheit:      -100,
			expectedCelcius: -73.33,
		},
		{
			name:            "positive value test",
			fahrenheit:      1298,
			expectedCelcius: 703.33,
		},
		{
			name:            "decimal value test",
			fahrenheit:      129.4566,
			expectedCelcius: 54.14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fahrenheit Conversion = &Temperature{
				Value: tt.fahrenheit,
			}
			result := fahrenheit.FahrenheitToCelcius(2)
			if result != tt.expectedCelcius {
				t.Errorf("conversion from fahrenheit to celcius failed: FahrenheitToCelcius(%v), expected celcius: %v", tt.fahrenheit, result)
			}
		})
	}
}
