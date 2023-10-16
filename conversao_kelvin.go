package main

import "fmt"

func main() {

	tempKelvin := 313
	fmt.Printf("%d Kelvin to Celsius is: %d", tempKelvin, converterParaCelsius(tempKelvin))
}

func converterParaCelsius(tempKelvin int) int {
	return tempKelvin - 273
}
