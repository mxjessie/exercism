// Package leap is a leap year figurer-outter.
package leap

// testVersion is, I guess, the version of the tests to run.
// I wish I knew what tool was making the Exercism bot comments.
// OH WELL. *iterates every 10 seconds*
const testVersion = 3

// isLeapYear is stolen from my rust version of this.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
