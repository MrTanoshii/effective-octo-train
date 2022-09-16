package lib

import (
	"fmt"
)

func LengthConverter(value float64, conversion_rate float64) float64 {
	result := value * conversion_rate

	fmt.Printf("Value: %v -> %v", value, result)
	return result
}
