package main
import (
        "fmt"
        "math"
        "regexp"
        "strconv"
        "strings"
)


func op_eval (a float64, b float64, op string) float64 {
        switch op {
        default: return b
        case  "+" : return a + b
        case  "-" : return a - b
        case  "*" : return a*b
        case  "/" : return a/b
        case  "^" : return math.Pow(a, b)
        }
}

func main() {

        str := "2*cos+2"
        re := []string{"\\+", "\\-", "\\*", "\\/", "\\^", "\\d", "[a-z]+"}
        rex := regexp.MustCompile("\\d")
        token := rex.FindString(str)

        if x, err := strconv.ParseFloat(token, 64) ; err == nil {
                fmt.Printf("%.2f \n", x*math.Cos(0))
                str = strings.Replace(str, token, "", 1)
                fmt.Println(str)
        }

}
