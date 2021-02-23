package main

import (
        "fmt"
        "runtime"
        "math"
)

func Echo(s []uint64, n uint64) {
        if n == 0 {
                fmt.Println(len(s))
                return
        }
        fmt.Println(s[n-1])
        Echo(s, n-1)
}

func Fcn(x float64) float64 {
        return x*x + 5
}

func DyDt(t float64, y float64) float64 {
        return y*t*t*t
}

func main() {
        x := math.Exp(4)
        xRK := RungeKutta(DyDt, 2, 0, 1, Mag(-5))
        xM := Midpoint(DyDt, 2, 0, 1, Mag(-5))
        xH := Heun(DyDt, 2, 0, 1, Mag(-5))
        xE := Euler(DyDt, 2, 0, 1, Mag(-5))
        fmt.Printf("%.12f\n", math.Abs(x - xM))
        fmt.Printf("%.12f\n", math.Abs(x - xH))
        fmt.Printf("%.12f\n", math.Abs(x - xE))
        fmt.Printf("%.12f\n", math.Abs(xH - xM))
        fmt.Println(runtime.GOMAXPROCS(0))
        fmt.Printf("%.64f\n", MachineErr())
        fmt.Printf("%.64e\n", MachineErr())
        fmt.Printf("%.64f\n", math.Log10(MachineErr()))
        fmt.Println(xRK)
        fmt.Println(math.Exp(4))
        fmt.Printf("%.12f\n", math.Abs(x - xRK))
}
