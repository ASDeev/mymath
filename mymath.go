package mymath

import "math"

// Abs возвращает абсолютное значение числа.
func Abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

// Max возвращает наибольшее число из двух.
func Max(x, y float64) float64 {
    if x > y {
        return x
    }
    return y
}

// Sqrt возвращает квадратный корень из числа.
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

// Sin возвращает синус угла в радианах.
func Sin(x float64) float64 {
    return math.Sin(x)
}

// Cos возвращает косинус угла в радианах.
func Cos(x float64) float64 {
    return math.Cos(x)
}

// Yn возвращает натуральное логарифмическое значение Бесселя функции первого рода.
func Yn(n int, x float64) float64 {
    return math.Yn(n, x)
}
