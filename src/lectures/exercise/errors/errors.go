//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"strconv"
	"strings"
)

type Time struct {
	hour int
	minute int
	second int
}

func ParseTime(timeString string) (Time, error) {
	components := strings.Split(timeString, ":")
	if (len(components) != 3) {
		return Time{}, errors.New("Time must have 3 components separated by a number")
	} else {
		hour, hourError := strconv.ParseInt(components[0], 10, 32)
		minute, minuteError := strconv.ParseInt(components[1], 10, 32)
		second, secondError := strconv.ParseInt(components[2], 10, 32)
		if (hourError != nil || minuteError != nil || secondError != nil) {
			return Time{}, errors.New("All time components must be numerical values")
		} else {
			return Time{int(hour), int(minute), int(second)}, nil
		}
	}
}