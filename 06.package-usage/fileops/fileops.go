package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// How ReadFile works here:
// - ReadFile reads a file, and convert it to slice of byte[] ( data var here)
// - Then those slice byte[] can be converted to string
// - Then convert that string to float64
func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // get slice of byte[]

	if err != nil {
		return 0, errors.New("failed to find balance in file")
	}

	fileText := string(data)                       // convert slice of byte[] to string
	value, err := strconv.ParseFloat(fileText, 64) // convert string to float

	if err != nil {
		return 0, errors.New("failed to parse stored balance value")
	}

	return value, nil
}

func WriteBalanceToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)             // convert float64 to string
	valueInBytes := []byte(valueText)          // convert string to byte[] slice
	os.WriteFile(fileName, valueInBytes, 0644) // write to file
}
