package lib

import (
	"github.com/MrTanoshii/effective-octo-train/Go/unit_converter/lib/constants"
)

func UnitToConversionRate(unit1 string, unit2 string) float64 {
	foundUnit1 := false
	foundUnit2 := false

	conversionRate := 0.0
	rate1 := 0.0
	rate2 := 0.0

	for k, v := range constants.UnitsLength {
		if !foundUnit1 && k == unit1 {
			rate1 = v
			foundUnit1 = true
		}
		if !foundUnit2 && k == unit2 {
			rate2 = v
			foundUnit2 = true
		}

		// Both units found
		if foundUnit1 && foundUnit2 {
			conversionRate = rate1 / rate2
			break
		}
	}
	return conversionRate
}
