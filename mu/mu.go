package mu
import ( "math" )

// ODEs
func RungeKutta(dy func(float64, float64) float64, tf float64, t float64, y float64, h float64) float64 {
        if tf <= t {
                return y
        }
        rk := []float64{h*dy(t, y)}
        for i := 1 ; i < 3 ; i++ {
                rk = append(rk, h*dy(t + h/2, y + rk[i-1]/2))
        }
        rk = append(rk, h*dy(t + h, y + rk[len(rk)-1]))
        y += (rk[0] + 2*rk[1] + 2*rk[2] + rk[3])/6
        t += h
        return RungeKutta(dy, tf, t, y, h)
}

func Heun(dy func(float64, float64) float64, tf float64, t float64, y float64, h float64) float64 {
        if tf <= t {
                return y
        }
        y += h/2*(dy(t, y) + dy(t + h, y + h*dy(t, y)))
        t += h
        return Heun(dy, tf, t, y, h)
}

func Midpoint(dy func(float64, float64) float64, tf float64, t float64, y float64, h float64) float64 {
        if tf <= t {
                return y
        }
        y += h*dy(t + h/2, y + h/2*dy(t, y))
        t += h
        return Midpoint(dy, tf, t, y, h)
}

func Euler(dy func(float64, float64) float64, tf float64, t float64, y float64, h float64) float64 {
        // should be used for comparison/education/curiosity only
        for t < tf {
                t += h
                y += h*dy(t, y)
        }
        return y
}

// Integrals
func Trapezoid(f func(float64) float64, a float64, b float64, n int) float64 {
        h := (b - a)/float64(n)
        var sum float64 = 0
        x := []float64{}
        for i := 0 ; i < n + 1 ; i++ {
                x = append(x, a + h*float64(i))
        }
        for i := 1 ; i < n ; i++ {
                sum += f(x[i])
        }
        return h*(sum + (f(a) + f(b))/2)
}

func Simpson(f func(float64) float64, a float64, b float64, n int) float64 {
        h := (b - a)/float64(2*n)
        x := []float64{}
        sum := []float64{0, 0}
        for i := 0 ; i < 2*n ; i++ {
                x = append(x, a + h*float64(i))
        }
        for i := 0 ; i < n ; i++ {
                sum[0] += f(x[2*i])
        }
        for i := 1 ; i < n ; i++ {
                sum[1] += f(x[2*i+1])
        }
        return h/3*(f(a) + f(b) + 4*sum[0] + 2*sum[1])
}

// First Order Differentiation
func ForwardDifff(f func(float64) float64, x float64, h float64) float64 {
        return (f(x + h) - f(x))/h
}

func CenteredDiff(f func(float64) float64, x float64, h float64) float64 {
        return (f(x + h) - f(x - h))/2*h
}

// Higher Order Differentiation
func NewtonDiff(f func(float64) float64, x float64, h float64, n uint64) float64 {
        if n == 0 {
                return f(x)
        }

        sum := 0.0
        var k uint64 = 0
        for k <= n {
                sum += math.Pow(-1, float64(k) + float64(n))*Binomial(n, k)*f(x + float64(k)*h)
                k++
        }
        return sum
}

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
        return math.Log2(x)/math.Log2(b)
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

func Fac(n uint64) uint64 {
        if n == 0 {
                return 1
        }
        if n == 1 {
                return n
        }
        return n*Fac(n-1)
}

func Binomial(n uint64, k uint64) float64 {
        return float64(Fac(n))/(float64(Fac(k)*Fac(n-k)))
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
