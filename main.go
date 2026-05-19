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
	log.Print("Listening on http://localhost:3000/...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
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
	// TODO:
}

func toMeters(unit string) float64 {
	// TODO:
}

// Temperature conversions
func convertTemperature(value float64, from, to string) float64 {
	// TODO:
}

// Area conversions
func convertArea(value float64, from, to string) float64 {
	// TODO:
}

func toSquareMeters(unit string) float64 {
	// TODO:
}

// Volume conversions
func convertVolume(value float64, from, to string) float64 {
	// TODO:
}

func toLiters(unit string) float64 {
	// TODO:
}

// Weight conversions
func convertWeight(value float64, from, to string) float64 {
	// TODO:
}

func toKilograms(unit string) float64 {
	// TODO:
}

// Time conversions
func convertTime(value float64, from, to string) float64 {
	// TODO:
}

func toSeconds(unit string) float64 {
	// TODO:
}
