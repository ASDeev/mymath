package mymath

import "math"

// Sqrt вычисляет квадратный корень из числа и возвращает результат в виде int
func Sqrt(x float64) int {
    return int(math.Sqrt(x))
}

// Abs возвращает абсолютное значение числа
func Abs(x int) int {
    return int(math.Abs(float64(x)))
}

// Max возвращает максимальное значение из двух чисел
func Max(x, y int) int {
    return int(math.Max(float64(x), float64(y)))
}

// Min возвращает минимальное значение из двух чисел
func Min(x, y int) int {
    return int(math.Min(float64(x), float64(y)))
}

// Round возвращает ближайшее целое число к данному числу с плавающей запятой
func Round(x float64) int {
    return int(math.Round(x))
}
