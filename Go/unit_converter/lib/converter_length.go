package lib

import (
	"fmt"

	"github.com/MrTanoshii/effective-octo-train/Go/unit_converter/lib/constants"
)

func UnitToConversionRate(unit1 string, unit2 string) float64 {
	fmt.Printf("Unit: %v -> %v", unit1, unit2)
	for k, v := range constants.UnitsLength {
		fmt.Println(k, v)
	}
}

func LengthConverter(value float64, conversion_rate float64) float64 {
	result := value * conversion_rate

	fmt.Printf("Value: %v -> %v", value, result)
	return result
}
