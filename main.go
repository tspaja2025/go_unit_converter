package main

import (
	"log"
	"net/http"
)

func main() {
	startServer()
}

// Unit
type Length struct {
	Meter      float64
	Kilometer  float64
	Centimeter float64
	Millimeter float64
	Micrometer float64
	Nanometer  float64
	Mile       float64
	Yard       float64
	Foot       float64
	Inch       float64
	LightYear  float64
}

type Temperature struct {
	Celcius    float64
	Kelvin     float64
	Fahrenheit float64
}

type Area struct {
	SquareMeter      float64
	SquareKilometer  float64
	SquareCentimeter float64
	SquareMillimeter float64
	SquareMicrometer float64
	Hectare          float64
	SquareMile       float64
	SquareYard       float64
	SquareFoot       float64
	SquareInch       float64
	Acre             float64
}

type Volume struct {
	CubicMeter         float64
	CubicKilometer     float64
	CubicCentimeter    float64
	CubicMillimeter    float64
	Liter              float64
	Milliliter         float64
	USGallon           float64
	USQuart            float64
	USPint             float64
	USCup              float64
	USFluidOunce       float64
	USTableSpoon       float64
	USTeaSpoon         float64
	ImperialGallon     float64
	ImperialQuart      float64
	ImperialPint       float64
	ImperialFluidOunce float64
	ImperialTableSpoon float64
	ImperialTeaSpoon   float64
	CubicMile          float64
	CubicYard          float64
	CubicFoot          float64
	CubicInch          float64
}

type Weight struct {
	Kilogram       float64
	Gram           float64
	Milligram      float64
	MetricTon      float64
	LongTon        float64
	ShortTon       float64
	Pound          float64
	Ounce          float64
	Carrat         float64
	AtomicMassUnit float64
}

type Time struct {
	Second      float64
	Millisecond float64
	Microsecond float64
	Nanosecond  float64
	Picosecond  float64
	Minute      float64
	Hour        float64
	Day         float64
	Week        float64
	Month       float64
	Year        float64
}

// Server
func startServer() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
