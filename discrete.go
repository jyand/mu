package main

func IntDiv(a int, b int) int {
        return (a - a%b)/b
}

func Primo(n uint64) bool {
        bl := true
        var i uint64 = 2
        for i < n {
                if n%i == 0 {
                        bl = false
                        break
                }
                i++
        }
        return bl
}

func Primes(n uint64) []uint64 {
        p := []uint64{}
        var im func(uint64)
        im = func(k uint64) {
                if k > 0 {
                        if Primo(k) {
                                p = append(p, k)
                        }
                        im(k-1)
                } else {
                        return
                }
        }
        im(n)
        return p
}

func Composites(n uint64) []uint64 {
        p := []uint64{}
        var i uint64 = 4
        for i <= n {
                if !Primo(uint64(i)) {
                        p = append(p, i)
                }
                i++
        }
        return p
}
