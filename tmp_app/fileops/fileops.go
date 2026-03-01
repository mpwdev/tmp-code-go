package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		//fmt.Println("Error reading balance file, use default balance:", defaultBalance)
		return defaultValue, errors.New("error reading file")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return defaultValue, errors.New("error parsing stored value")
	}
	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
