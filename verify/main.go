package main

import (
        "fmt"
        "runtime"
        "math"
        "mu/mu"
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
        xRK := mu.RungeKutta(DyDt, 2, 0, 1, mu.Mag(-5))
        xM := mu.Midpoint(DyDt, 2, 0, 1, mu.Mag(-5))
        xH := mu.Heun(DyDt, 2, 0, 1, mu.Mag(-5))
        xE := mu.Euler(DyDt, 2, 0, 1, mu.Mag(-5))
        fmt.Printf("%.12f\n", math.Abs(x - xM))
        fmt.Printf("%.12f\n", math.Abs(x - xH))
        fmt.Printf("%.12f\n", math.Abs(x - xE))
        fmt.Printf("%.12f\n", math.Abs(xH - xM))
        fmt.Println(runtime.GOMAXPROCS(0))
        fmt.Printf("%.64f\n", mu.MachineErr())
        fmt.Printf("%.64e\n", mu.MachineErr())
        fmt.Printf("%.64f\n", math.Log10(mu.MachineErr()))
        fmt.Println(xRK)
        fmt.Println(math.Exp(4))
        fmt.Printf("%.12f\n", math.Abs(x - xRK))
        fmt.Println(mu.Fac(24))
        fmt.Println(mu.NewtonDiff(Fcn, 1, mu.Mag(2), 1))
        fmt.Println(mu.LogBase(5, 625))
}
