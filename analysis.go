package main

func RungeKutta(dy func(float64, float64) float64, tf float64, t float64, y float64, h float64) float64 {
        if tf <= t {
                return y
        }
        y += h*dy(t + h/2, y + h/2*dy(t, y))
        t += h
        return RungeKutta(dy, tf, t, y, h)
}

func Eulers(dy func(float64, float64) float64, tf float64, ti float64, yi float64, h float64) float64 {
        t := ti
        y := yi
        for t < tf {
                t += h
                y += h*dy(t, y)
        }
        return y
}

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

func Simpsons(f func(float64) float64, a float64, b float64, n int) float64 {
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
