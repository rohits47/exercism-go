package space

// EarthYearSeconds is the number of seconds in a calendar year on earth.
// This is equivalent to 365.25 days.
const EarthYearSeconds = 31557600

// Age contains the seconds in the age and the associated planet.
type Age struct {
	seconds int
	planet  string
}
