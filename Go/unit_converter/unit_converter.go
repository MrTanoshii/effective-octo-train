package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MrTanoshii/effective-octo-train/Go/unit_converter/lib"
)

func main() {
	var command string
	var value float64

	// Flags declaration
	flag.StringVar(&command, "c", "length m mi", "Specify command. Default is length meter to miles.")
	flag.Float64Var(&value, "v", 1, "Specify value. Default is 1.")

	flag.Parse()

	fmt.Println("~~ Unit Converter ~~")
	fmt.Println()

	fmt.Println("Received command:", command, "with values: ", value)
	fmt.Println()

	// Accept only 3 arguments in the command
	commandArray := strings.Split(command, " ")
	if len(commandArray) != 3 {
		return
	}

	switch commandArray[0] {
	case "length":
		fmt.Println("Converter check: [Found]")
	default:
		fmt.Println("Invalid command. Please try again.")
		return
	}

	var conversionRate float64 = lib.UnitToConversionRate(commandArray[1], commandArray[2])
	if conversionRate == 0.0 {
		return
	} else {
		fmt.Println("Conversion rate from", commandArray[1], "to", commandArray[2], "is", conversionRate)
	}
}
