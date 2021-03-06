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
