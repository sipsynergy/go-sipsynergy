package utils

import "strconv"

func ConvertUintToString(value uint) string {
    return strconv.Itoa(int(value))
}

func ConvertStringToUint(value string) (uint, error) {
    valueInt, err := strconv.ParseUint(value, 10, 64)
    return uint(valueInt), err
}

func ConvertStringToInt(value string) (int, error) {
    return strconv.Atoi(value)
}

func ConvertIntToString(value int) string {
    return strconv.Itoa(value)
}

func ConvertFloatToString(value float64) string {
    return strconv.FormatFloat(value, 'g', 1, 64)
}

func ConvertStringToFloat(value string, precision int) (float64, error) {
    return strconv.ParseFloat(value, precision)
}
