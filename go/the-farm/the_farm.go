package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(c FodderCalculator, cows int) (float64, error) {
	f, err := c.FatteningFactor()
	if err != nil {
		return 0, err
	}
	a, err := c.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	return a * f / float64(cows), nil
}

func ValidateInputAndDivideFood(c FodderCalculator, cows int) (float64, error) {
	err := ValidateNumberOfCows(cows)
	if err != nil {
		return 0, errors.New("invalid number of cows")
	}
	a, err := DivideFood(c, cows)
	if err != nil {
		return 0, err
	}
	return a, nil
}

type InvalidNumberOfCowsError struct {
	count int
	message string
}

func (e *InvalidNumberOfCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.count, e.message)
}

func ValidateNumberOfCows(c int) error {
	if c == 0 {
		return &InvalidNumberOfCowsError{c, "no cows don't need food"}
	} else if c < 0 {
		return &InvalidNumberOfCowsError{c, "there are no negative cows"}
	} else {
		return nil
	}
}
