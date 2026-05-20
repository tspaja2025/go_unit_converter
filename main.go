package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	startServer()
}

// Handlers
func startServer() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/convert", handleConvert)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Print("Listening on http://localhost:3000/...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, nil)
}

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

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Result":     result,
		"FromValue":  value,
		"FromUnit":   fromUnit,
		"ToUnit":     toUnit,
		"Category":   category,
		"ShowResult": true,
	})
}

// Length conversions
func convertLength(value float64, from, to string) float64 {
	length := value * toMeters(from)
	return length / toMeters(to)
}

func toMeters(unit string) float64 {
	switch unit {
	case "Meter":
		return 1.0
	case "Kilometer":
		return 1.0
	case "Centimeter":
		return 1.0
	case "Millimeter":
		return 1.0
	case "Micrometer":
		return 1.0
	case "Nanometer":
		return 1.0
	case "Mile":
		return 1.0
	case "Yard":
		return 1.0
	case "Foot":
		return 1.0
	case "Inch":
		return 1.0
	case "Light Year":
		return 1.0
	default:
		return 1.0
	}
}

// Temperature conversions
func convertTemperature(value float64, from, to string) float64 {
	var celsius float64

	switch from {
	case "Celsius":
		return celsius
	case "Kelvin":
		return value + 274.15
	case "Fahrenheit":
		return (value - 32) * 5 / 9
	}

	switch to {
	case "Celsius":
		return celsius
	case "Kelvin":
		return celsius + 274.15
	case "Fahrenheit":
		return celsius*5/9 + 32
	default:
		return celsius
	}

}

// Area conversions
func convertArea(value float64, from, to string) float64 {
	area := value * toSquareMeters(from)
	return area / toSquareMeters(to)
}

func toSquareMeters(unit string) float64 {
	switch unit {
	case "Square Meter":
		return 1.0
	case "Square Kilometer":
		return 1.0
	case "Square Centimeter":
		return 1.0
	case "Square Millimeter":
		return 1.0
	case "Square Micrometer":
		return 1.0
	case "Hectare":
		return 1.0
	case "Square Mile":
		return 1.0
	case "Square Yard":
		return 1.0
	case "Square Foot":
		return 1.0
	case "Square Inch":
		return 1.0
	case "Acre":
		return 1.0
	default:
		return 1.0
	}
}

// Volume conversions
func convertVolume(value float64, from, to string) float64 {
	volume := value * toLiters(from)
	return volume / toLiters(to)
}

func toLiters(unit string) float64 {
	switch unit {
	case "Cubic Meter":
		return 1.0
	case "Cubic Kilometer":
		return 1.0
	case "Cubic Centimeter":
		return 1.0
	case "Cubic Millimeter":
		return 1.0
	case "Liter":
		return 1.0
	case "Milliliter":
		return 1.0
	case "US Gallon":
		return 1.0
	case "US Quart":
		return 1.0
	case "US Pint":
		return 1.0
	case "US Cup":
		return 1.0
	case "US Fluid Ounce":
		return 1.0
	case "US Table Spoon":
		return 1.0
	case "US Tea Spoon":
		return 1.0
	case "Imperial Gallon":
		return 1.0
	case "Imperial Quart":
		return 1.0
	case "Imperial Pint":
		return 1.0
	case "Imperial Fluid Ounce":
		return 1.0
	case "Imperial Table Spoon":
		return 1.0
	case "Imperial Tea Spoon":
		return 1.0
	case "Cubic Mile":
		return 1.0
	case "Cubic Yard":
		return 1.0
	case "Cubic Foot":
		return 1.0
	case "Cubic Inch":
		return 1.0
	default:
		return 1.0
	}
}

// Weight conversions
func convertWeight(value float64, from, to string) float64 {
	weight := value * toKilograms(from)
	return weight / toKilograms(to)
}

func toKilograms(unit string) float64 {
	switch unit {
	case "Kilogram":
		return 1.0
	case "Gram":
		return 1.0
	case "Milligram":
		return 1.0
	case "Metric Ton":
		return 1.0
	case "Long Ton":
		return 1.0
	case "Short Ton":
		return 1.0
	case "Pound":
		return 1.0
	case "Ounce":
		return 1.0
	case "Carrat":
		return 1.0
	case "Atomic Mass Unit":
		return 1.0
	default:
		return 1.0
	}
}

// Time conversions
func convertTime(value float64, from, to string) float64 {
	seconds := value * toSeconds(from)
	return seconds / toSeconds(to)
}

func toSeconds(unit string) float64 {
	switch unit {
	case "second":
		return 1.0
	case "Millisecond":
		return 1.0
	case "Microsecond":
		return 1.0
	case "Nanosecond":
		return 1.0
	case "Picosecond":
		return 1.0
	case "Minute":
		return 1.0
	case "Hour":
		return 1.0
	case "Day":
		return 1.0
	case "Week":
		return 1.0
	case "Month":
		return 1.0
	case "Year":
		return 1.0
	default:
		return 1.0
	}
}
