package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	converted := make([]float64, len(strings))
	for i, v := range strings {
		floatPrice, err := strconv.ParseFloat(v, 64)
		converted[i] = floatPrice

		if err != nil {
			return nil, errors.New("failed to convert strings to floats")
		}
	}
	return converted, nil
}
