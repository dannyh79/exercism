// Package leap implements leap-year-related utility functions
package leap

/*
Determine whether a year is a leap year (in the Gregorian calendar):

- In every year that is evenly divisible by 4.

- Unless the year is evenly divisible by 100, in which case it's only a leap year if the year is also evenly divisible by 400.
*/
func IsLeapYear(year int) bool {
	return (year % 100 != 0 && year % 4 == 0) || (year % 400 == 0)
}
