package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64,fileName string){
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName,[]byte(valueText),0644)
}

func GetFloatFromfile(fileName string)(float64, error){
	data, err := os.ReadFile(fileName)

	if err != nil {
 	return  1000 ,errors.New("failed to parse file")
	}
    valueText := string(data)
	value, err := strconv.ParseFloat(valueText,64)
	
	if err!= nil {
		return 1000.0 , errors.New("failed to parse file")
	}

	return value, nil 
}