package main

import (
        "fmt"
        "runtime"
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
        fmt.Println(RungeKutta(DyDt, 2, 0, 1, 0.00001))
        fmt.Println(Heun(DyDt, 2, 0, 1, 0.00001))
        fmt.Println(Euler(DyDt, 2, 0, 1, 0.00001))
        fmt.Println(runtime.GOMAXPROCS(0))
        /*fmt.Println(IntDiv(8, 3))
        fmt.Println(Primo(13))
        fmt.Println(Primo(14))
        l := Primes(499)
        Echo(l, uint64(len(l)))
        fmt.Println()
        l = Composites(499)
        Echo(l, uint64(len(l)))
        fmt.Println(IntSqrt(7))
        fmt.Println(IntSqrt(30))
        fmt.Println(Totient(15))
        fmt.Println(Trapezoid(Fcn, 0, 10, 100000000))
        fmt.Println(Simpson(Fcn, 0, 10, 100000000))*/
}
