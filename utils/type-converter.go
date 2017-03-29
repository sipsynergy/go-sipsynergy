package utils

import "strconv"

func ConvertUintToString(value uint) string {
    return strconv.Itoa(int(value))
}

func ConvertStringToUint(value string) (uint, error) {
    uint64value, err := strconv.ParseUint(value, 10, 64)
    return uint(uint64value), err
}

func ConvertStringToInt(value string) int {
    i, _ := strconv.Atoi(value)
    return i
}

func ConvertIntToString(value int) string {
    return strconv.Itoa(value)
}

func ConvertFloatToString(value float64) string {
    return strconv.FormatFloat(value, 'g', 1, 64)
}

func ConvertStringToFloat(value string, precision int) float64 {
    f, _ := strconv.ParseFloat(value, precision)
    return f
}