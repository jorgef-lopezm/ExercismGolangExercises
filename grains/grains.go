package grains

import "errors"
import "math"

func Square(square int) (uint64, error) {

	if square <= 0 || square > 64 {
		return 0, errors.New("Value is out of range.")
	}

	return uint64(math.Exp2(float64(square - 1))), nil
}

func Total() uint64 {
	return math.MaxUint64
}