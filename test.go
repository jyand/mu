/* mu.go 
* main part of the application that runs in a shell and processes CLI arguments 
*/
package main
import (
        "fmt"
        "os"
        "math"
//        "encoding/csv"
//       "encoding/json"
//        "regexp"
        "strconv"
)

func feval (f, x float64) float64 {
        /* convert commandline syntax to math pkg syntax */
        f = map[string]func(float64) float64 {
                /* single-variable, real-valued, differentiable functions */
                "cos": math.Cos,
                "sin": math.Sin,
                "e^": math.Exp,
                "log": math.Log2,
        }
        return f(x)
}

func main() {

        // testing/debugging
        if x, err := strconv.ParseFloat(os.Args[2], 64) ; err == nil {
                fmt.Printf("%.2f \n", feval(f[os.Args[1]], x))
        }
}
