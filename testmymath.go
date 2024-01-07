package testmymath

import "math"

func TestMySqrt(x float64) float64 {
	return math.Sqrt(x)
}

func TestMyCeil(x float64) float64 {
	return math.Ceil(x)
}

func TestMyFloor(x float64) float64 {
	return math.Floor(x)
}

func TestMyPow(x, y float64) float64 {
	return math.Pow(x, y)
}

func TestMyMax(x, y float64) float64 {
	return math.Max(x, y)
}

func TestMyMin(x, y float64) float64 {
	return math.Min(x, y)
}
