package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

var ZeroCowsError = errors.New("division by zero")
var NegativeFodderError = errors.New("negative fodder")

func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, error := weightFodder.FodderAmount()

	if hasError(fodderAmount, cows, error) {
		return handleError(fodderAmount, float64(cows), error)
	}

	return fodderAmount / float64(cows), nil
}

func hasError(fodderAmount float64, cows int, err error) bool {
	return fodderAmount < 0.0 || cows <= 0 || err != nil
}

func handleError(fodderAmount, cows float64, err error) (float64, error) {
	if err == ErrScaleMalfunction {
		return handleErrScaleMalfunction(fodderAmount, cows)
	}

	if err != nil {
		return 0.0, err
	}

	if fodderAmount < 0.0 {
		return 0.0, NegativeFodderError
	}

	if cows == 0.0 {
		return 0.0, ZeroCowsError
	}

	return 0.0, &SillyNephewError{cows: int(cows)}
}

func handleErrScaleMalfunction(fodderAmount, cows float64) (float64, error) {
	if fodderAmount <= 0.0 {
		return 0.0, NegativeFodderError
	}

	fodderPerCow := (2.0 * fodderAmount) / cows

	return fodderPerCow, nil
}
