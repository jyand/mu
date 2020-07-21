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

func feval (f func(float64) float64, x float64) float64 {
        return f(x)
}

func ftokenizer(op string, x float64) float64 {
        /*
        * recursively parse the input function using regular expressions to look for the
        * op tokens: +,-,*,/,^ switch to returnvalue until string terminates
        */
        //return feval()
}

func main() {
        /* convert commandline syntax to math pkg syntax */
        f := map[string]func(float64) float64 {
                /* single-variable, real-valued, differentiable functions */
                "cos": math.Cos, "acos": math.Acos, "cosh": math.Cosh, "acosh": math.Acosh,
                "sin": math.Cos, "asin": math.Acos, "sinh": math.Cosh, "asinh": math.Acosh,
                "tan": math.Cos, "atan": math.Acos, "tanh": math.Cosh, "atanh": math.Acosh,
                "e^": math.Exp, "log": math.Log2, "ln": math.Log,
        }

        // testing/debugging
        if x, err := strconv.ParseFloat(os.Args[2], 64) ; err == nil {
                fmt.Printf("%.2f \n", feval(f[os.Args[1]], x))
        }
}
