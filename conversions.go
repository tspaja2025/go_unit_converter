package main

import (
	"log"
	"net/http"
	"strconv"
)

func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	category := r.FormValue("category")
	fromUnit := r.FormValue("fromUnit")
	toUnit := r.FormValue("toUnit")
	valueStr := r.FormValue("value")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	var result float64
	switch category {
	case "length":
		result = convertLength(value, fromUnit, toUnit)
	case "temperature":
		result = convertTemperature(value, fromUnit, toUnit)
	case "area":
		result = convertArea(value, fromUnit, toUnit)
	case "volume":
		result = convertVolume(value, fromUnit, toUnit)
	case "weight":
		result = convertWeight(value, fromUnit, toUnit)
	case "time":
		result = convertTime(value, fromUnit, toUnit)
	}

	data := PageData{
		Result:     result,
		FromValue:  value,
		FromUnit:   fromUnit,
		ToUnit:     toUnit,
		Category:   category,
		ShowResult: true,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Conversion helper
func convert(value float64, fromUnit, toUnit string, conversionMap map[string]float64) float64 {
	fromFactor, fromExists := conversionMap[fromUnit]
	toFactor, toExists := conversionMap[toUnit]

	if !fromExists || !toExists {
		return value // Return original if units not found
	}

	// Convert to base unit first, then to target unit
	baseValue := value * fromFactor
	return baseValue / toFactor
}

// Length conversions
func convertLength(value float64, from, to string) float64 {
	return convert(value, from, to, lengthUnits)
}

// Temperature conversions
func convertTemperature(value float64, from, to string) float64 {
	c := toCelsius[from](value)
	return fromCelsius[to](c)
}

// Area conversions
func convertArea(value float64, from, to string) float64 {
	return convert(value, from, to, areaUnits)
}

// Volume conversions
func convertVolume(value float64, from, to string) float64 {
	return convert(value, from, to, volumeUnits)
}

// Weight conversions
func convertWeight(value float64, from, to string) float64 {
	return convert(value, from, to, weightUnits)
}

// Time conversions
func convertTime(value float64, from, to string) float64 {
	return convert(value, from, to, timeUnits)
}
