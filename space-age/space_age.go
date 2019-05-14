package space

// EarthYearSeconds is the number of seconds in a calendar year on earth.
// This is equivalent to 365.25 days.
const EarthYearSeconds = 31557600

// EarthYearDays is the number of days that one earth year maps to on Earth
const EarthYearDays = 365.25

// MercuryYearDays is the number of days that one earth year maps to on Mercury
const MercuryYearDays = 0.2408467

// VenusYearDays is the number of days that one earth year maps to on Venus
const VenusYearDays = 0.61519726

// MarsYearDays is the number of days that one earth year maps to on Mars
const MarsYearDays = 1.8808158

// JupiterYearDays is the number of days that one earth year maps to on Jupiter
const JupiterYearDays = 11.862615

// SaturnYearDays is the number of days that one earth year maps to on Saturn
const SaturnYearDays = 29.447498

// UranusYearDays is the number of days that one earth year maps to on Uranus
const UranusYearDays = 84.016846

// NeptuneYearDays is the number of days that one earth year maps to on Neptune
const NeptuneYearDays = 164.79132

// Planet is a named planet in our solar system.
type Planet string

// Age gives the earth age in days from the given age on the given planet
func Age(seconds float64, planet Planet) float64 {
	var convertedAge float64
	convertedAge = seconds / EarthYearSeconds
	if planet == "Mercury" {
		convertedAge = convertedAge / MercuryYearDays
	}
	if planet == "Venus" {
		convertedAge = convertedAge / VenusYearDays
	}
	if planet == "Mars" {
		convertedAge = convertedAge / MarsYearDays
	}
	if planet == "Jupiter" {
		convertedAge = convertedAge / JupiterYearDays
	}
	if planet == "Saturn" {
		convertedAge = convertedAge / SaturnYearDays
	}
	if planet == "Uranus" {
		convertedAge = convertedAge / UranusYearDays
	}
	if planet == "Neptune" {
		convertedAge = convertedAge / NeptuneYearDays
	}
	return convertedAge
}
