package main

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
	"Carat":            0.0002,
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

var toCelsius = map[string]func(float64) float64{
	"Celsius":    func(f float64) float64 { return f },
	"Kelvin":     func(f float64) float64 { return f - 273.15 },
	"Fahrenheit": func(f float64) float64 { return (f - 32) * 5 / 9 },
}

var fromCelsius = map[string]func(float64) float64{
	"Celcius":    func(f float64) float64 { return f },
	"Kelvin":     func(f float64) float64 { return f + 273.15 },
	"Fahrenheit": func(f float64) float64 { return f*9/5 + 32 },
}
