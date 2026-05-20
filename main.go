package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageData struct {
	Result     float64
	FromValue  float64
	FromUnit   string
	ToUnit     string
	Category   string
	ShowResult bool
}

func main() {
	startServer()
}

// Handlers
func startServer() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/convert", handleConvert)
	fs := http.FileServer(http.Dir("static"))                 // for css
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // for css
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
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, PageData{
		Result:     result,
		FromValue:  value,
		FromUnit:   fromUnit,
		ToUnit:     toUnit,
		Category:   category,
		ShowResult: true,
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
		return 1000.0
	case "Centimeter":
		return 0.01
	case "Millimeter":
		return 0.001
	case "Micrometer":
		return 0.000001
	case "Nanometer":
		return 0.000000001
	case "Mile":
		return 1609.344
	case "Yard":
		return 0.9144
	case "Foot":
		return 0.3048
	case "Inch":
		return 0.0254
	case "Light Year":
		return 9.461e15
	default:
		return 1.0
	}
}

// Temperature conversions
func convertTemperature(value float64, from, to string) float64 {
	var celsius float64

	switch from {
	case "Celsius":
		celsius = value
	case "Kelvin":
		celsius = value - 273.15
	case "Fahrenheit":
		celsius = (value - 32) * 5 / 9
	}

	switch to {
	case "Celsius":
		return celsius
	case "Kelvin":
		return celsius + 273.15
	case "Fahrenheit":
		return celsius*9/5 + 32
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
		return 1000000.0
	case "Square Centimeter":
		return 0.0001
	case "Square Millimeter":
		return 0.000001
	case "Square Micrometer":
		return 0.000000000001
	case "Hectare":
		return 10000.0
	case "Square Mile":
		return 2589988.11
	case "Square Yard":
		return 0.836127
	case "Square Foot":
		return 0.092903
	case "Square Inch":
		return 0.00064516
	case "Acre":
		return 4046.86
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
		return 1000.0
	case "Cubic Kilometer":
		return 1e12
	case "Cubic Centimeter":
		return 0.001
	case "Cubic Millimeter":
		return 0.000001
	case "Liter":
		return 1.0
	case "Milliliter":
		return 0.001
	case "US Gallon":
		return 3.78541
	case "US Quart":
		return 0.946353
	case "US Pint":
		return 0.473176
	case "US Cup":
		return 0.24
	case "US Fluid Ounce":
		return 0.0295735
	case "US Table Spoon":
		return 0.0147868
	case "US Tea Spoon":
		return 0.00492892
	case "Imperial Gallon":
		return 4.54609
	case "Imperial Quart":
		return 1.13652
	case "Imperial Pint":
		return 0.568261
	case "Imperial Fluid Ounce":
		return 0.0284131
	case "Imperial Table Spoon":
		return 0.0177582
	case "Imperial Tea Spoon":
		return 0.00591939
	case "Cubic Mile":
		return 4.16818e12
	case "Cubic Yard":
		return 764.555
	case "Cubic Foot":
		return 28.3168
	case "Cubic Inch":
		return 0.0163871
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
