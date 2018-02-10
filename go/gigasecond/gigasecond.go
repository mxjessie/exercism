// Package gigasecond provides addGigasecond for adding a gigasecond to a time.Time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds one (1) gigasecond to a time.Time.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1e9) * time.Second
	return t.Add(gigasecond)
}
