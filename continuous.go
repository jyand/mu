package main

import (
        "math"
)

func Mag(order int8) float64 {
        return math.Pow(10, float64(order))
}

func MachineErr() float64 {
        eps := 1.0
        for 1.0 + eps*0.5 > 1.0 {
                eps *= 0.5
        }
        return eps
}

func LogBase(b float64, x float64) float64 {
        return math.Log(x)/math.Log(b)
}

func Sigmoid(x float64) float64 {
        return (math.Expm1(x) + 1)/(math.Exp(x) + 1)
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
