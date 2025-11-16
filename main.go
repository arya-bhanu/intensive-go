package main

import (
	"fmt"
	"math"
)

func main() {
	var celcius Conversion = &Temperature{
		Value: 60,
	}

	var fahrenheit Conversion = &Temperature{
		Value: 80,
	}

	fmt.Println(celcius.CelciusToFahrenheit(3))
	fmt.Println(fahrenheit.FahrenheitToCelcius(1))
	fmt.Printf("fahrenheit test: %#v\n", fahrenheit)
}

type Temperature struct {
	Value float64
}

type Conversion interface {
	CelciusToFahrenheit(round int) float64
	FahrenheitToCelcius(round int) float64
}

func (t *Temperature) FahrenheitToCelcius(round int) float64 {
	result := (t.Value - 32) * float64(5) / float64(9)
	if round != 0 {
		return roundNumber(result, round)
	}
	return result
}

func (t *Temperature) CelciusToFahrenheit(round int) float64 {
	result := t.Value*(float64(9)/float64(5)) + 32
	if round != 0 {
		return roundNumber(result, round)
	}
	return result
}

func roundNumber(val float64, round int) float64 {
	pow := math.Pow10(round)
	return math.Round(val*pow) / pow
}
