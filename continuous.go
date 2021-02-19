package main

import (
        "math"
)

func LogBase(b float64, x float64) float64 {
        return math.Log2(x)/math.Log2(b)
}

func Step(x float64) float64 {
        if x < 0 {
                return 0
        } else {
                return 1
        }
}

func Signum(x float64) float64 {
        if x == 0 {
                return 0
        }
        if x < 0 {
                return -1
        } else {
                return 1
        }
}

func Ramp(x float64) float64 {
        if x < 0 {
                return 0
        } else {
                return x
        }
}
