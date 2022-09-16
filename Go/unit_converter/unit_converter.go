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

	fmt.Println()
	fmt.Println("~~ Unit Converter ~~")

	fmt.Println()
	fmt.Println("Received command:", command)
	fmt.Println("Received value:", value)

	// Accept only 3 arguments in the command
	commandArray := strings.Split(command, " ")
	if len(commandArray) != 3 {
		return
	}

	var conversionRate float64
	switch commandArray[0] {
	case "length":
		fmt.Println()
		fmt.Println("[Found] Converter")

		conversionRate = lib.UnitToConversionRate(commandArray[1], commandArray[2])
		if conversionRate == 0.0 {
			fmt.Println("[Invalid] Conversion units")
			return
		} else {
			fmt.Println("[Found] Conversion units (rate:", conversionRate, ")")
		}
	default:
		fmt.Println("[Invalid] Converter")
		return
	}

	if commandArray[0] == "length" {
		fmt.Println()
		displayResult(commandArray, value, conversionRate)
	}

	fmt.Println()
	fmt.Println("~~ Goodbye ~~")
	fmt.Println()
}

func displayResult(commandArray []string, value float64, conversionRate float64) {
	fmt.Println(value, commandArray[1], "=", lib.LengthConverter(value, conversionRate), commandArray[2])
}
