package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type PageData struct {
	Result     float64
	FromValue  float64
	FromUnit   string
	ToUnit     string
	Category   string
	ShowResult bool
}

var lengthUnits = map[string]float64{
	"Meter":      1.0,
	"Kilometer":  1000.0,
	"Centimeter": 0.01,
	"Millimeter": 0.001,
	"Micrometer": 0.000001,
	"Nanometer":  0.000000001,
	"Mile":       1609.344,
	"Yard":       0.9144,
	"Foot":       0.3048,
	"Inch":       0.0254,
	"Light Year": 9.461e15,
}

var areaUnits = map[string]float64{
	"Square Meter":      1.0,
	"Square Kilometer":  1000000.0,
	"Square Centimeter": 0.0001,
	"Square Millimeter": 0.000001,
	"Square Micrometer": 0.000000000001,
	"Hectare":           10000.0,
	"Square Mile":       2589988.11,
	"Square Yard":       0.836127,
	"Square Foot":       0.092903,
	"Square Inch":       0.00064516,
	"Acre":              4046.86,
}

var volumeUnits = map[string]float64{
	"Cubic Meter":          1000.0,
	"Cubic Kilometer":      1e12,
	"Cubic Centimeter":     0.001,
	"Cubic Millimeter":     0.000001,
	"Liter":                1.0,
	"Milliliter":           0.001,
	"US Gallon":            3.78541,
	"US Quart":             0.946353,
	"US Pint":              0.473176,
	"US Cup":               0.24,
	"US Fluid Ounce":       0.0295735,
	"US Table Spoon":       0.0147868,
	"US Teaspoon":          0.00492892,
	"Imperial Gallon":      4.54609,
	"Imperial Quart":       1.13652,
	"Imperial Pint":        0.568261,
	"Imperial Fluid Ounce": 0.0284131,
	"Imperial Table Spoon": 0.0177582,
	"Imperial Teaspoon":    0.00591939,
	"Cubic Mile":           4.16818e12,
	"Cubic Yard":           764.555,
	"Cubic Foot":           28.3168,
	"Cubic Inch":           0.0163871,
}

var weightUnits = map[string]float64{
	"Kilogram":         1.0,
	"Gram":             0.001,
	"Milligram":        0.000001,
	"Metric Ton":       1000.0,
	"Long Ton":         1016.05,
	"Short Ton":        907.185,
	"Pound":            0.453592,
	"Ounce":            0.0283495,
	"carat":            0.0002,
	"Atomic Mass Unit": 1.66054e-27,
}

var timeUnits = map[string]float64{
	"Second":      1.0,
	"Millisecond": 0.001,
	"Microsecond": 0.000001,
	"Nanosecond":  1e-9,
	"Picosecond":  1e-12,
	"Minute":      60.0,
	"Hour":        3600.0,
	"Day":         86400.0,
	"Week":        604800.0,
	"Month":       2628000.0,
	"Year":        31536000.0,
}

var tmpl *template.Template

var (
	tmplOnce sync.Once
	tmplLazy *template.Template
)

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

// Parse templates once when package initializes
func init() {
	var err error
	tmpl, err = template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}
	log.Println("Templates parsed successfully")
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
