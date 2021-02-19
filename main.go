package main

import (
        "fmt"
)

func Echo(s []uint64, n uint64) {
        if n == 0 {
                fmt.Println(len(s))
                return
        }
        fmt.Println(s[n-1])
        Echo(s, n-1)
}

func main() {
        /*fmt.Println(IntDiv(8, 3))
        fmt.Println(Primo(13))
        fmt.Println(Primo(14))*/
        l := Primes(499)
        Echo(l, uint64(len(l)))
        fmt.Println()
        l = Composites(499)
        Echo(l, uint64(len(l)))
}
