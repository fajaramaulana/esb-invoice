package helper

import "strconv"

func ConvertStringToInt(str string) (int, error) {
	stringConv, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return stringConv, nil
}

func MinLengthQueryParam(str string, min int) bool {
	if str != "" && len(str) < min {
		return true
	}
	return false
}
