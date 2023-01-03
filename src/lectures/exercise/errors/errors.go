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
	"fmt"
	"strconv"
	"strings"
)

type Timestamp struct {
	hour, minute, second int
}

func parse(s string) (Timestamp, error) {
	components := strings.Split(s, ":")
	if len(components) != 3 {
		return Timestamp{}, errors.New(fmt.Sprintf("%v cannot be parsed into timestamp", s))
	}
	convertedComponents := make([]int, 3)
	var component int
	var err error
	for i := 0; i < 3; i++ {
		component, err = strconv.Atoi(components[i])
		if err != nil {
			return Timestamp{}, errors.New(fmt.Sprintf("%v cannot be converted into a component", component))

		}

		convertedComponents[i] = component
	}

	isHourValid := (convertedComponents[0] < 24) && (convertedComponents[0] >= 0)
	isMinuteValid := (convertedComponents[1] < 60) && (convertedComponents[1] >= 0)
	isSecondValid := (convertedComponents[2] < 60) && (convertedComponents[2] >= 0)

	if !(isHourValid && isMinuteValid && isSecondValid) {
		return Timestamp{}, errors.New(fmt.Sprintf("%v has invalid timestamp values", s))
	}

	return Timestamp{
		hour:   convertedComponents[0],
		minute: convertedComponents[1],
		second: convertedComponents[2]}, nil
}
