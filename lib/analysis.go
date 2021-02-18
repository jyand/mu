/* lib/analysis.go 
* methods for solving ordinary differential equations and evaluating derivatives/integrals numerically
*/
package main
import (
        "math"
)

func rk (y_prime func(float64) float64, a float64, b float64, y_init float64, stepsize float64) float64 {
        return
}

func euler (y_prime func(float64) float64, a float64, b float64, y_init float64, stepsize float64) float64 {
        return
}

func trapezoid (f func(float64) float64, a float64, b float64, n int) float64 {
        return
}

func simpsons (f func(float64) float64, a float64, b float64, n int) float64 {
        return
}
