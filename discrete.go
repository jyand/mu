package main

func IntDiv(a uint64, b uint64) uint64 {
        return (a - a%b)/b
}

func IntSqrt(n uint64) uint64 {
        var im func(uint64, uint64) uint64
        im = func(k uint64, i uint64) uint64 {
                if k - IntDiv(k, i) <= i*i {
                        return i
                }
                return im(k, i+1)
        }
        return im(n, 2)
}

func Fac(n uint) uint {
        if n == 0 {
                return 1
        }
        if n == 1 {
                return n
        }
        return n*Fac(n-1)
}

func Euclidean(a uint64, b uint64) uint64 {
        if b == 0 {
                return a
        }
        return Euclidean(b, a%b)
}

func Totient(n uint64) uint64 {
        var im func(uint64, uint64) uint64
        im = func(k uint64, phi uint64) uint64 {
                if k < 1 {
                        return phi
                }
                if Euclidean(n, k) == 1{
                        phi++
                }
                return im(k-1, phi)
        }
        return im(n-1, 0)
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

//use this method for reals or just forget about it, type conversions are a pain
func EulersTotient(n uint64) uint64 {
        p := Primes(n)
        for i := 0 ; i < len(p) ; i++ {
                if n%p[i] == 0 {
                        n = n*(1 - 1/p[i])
                }
        }
        return n
}
