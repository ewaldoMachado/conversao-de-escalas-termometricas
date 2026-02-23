package main

import "fmt"

func main() {
	var tKelvin = 373.15
	var tCelcius = tKelvin - 273.15
	var atm = "1 atm"
	fmt.Printf("O ponto de ebulição da água é %g°C ao nível do mar, em condições normais de pressão atmosférica (%s)", tCelcius, atm)
}