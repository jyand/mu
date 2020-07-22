package main
import (
        "fmt"
//       "math"
        "regexp"
//        "strconv"
//        "strings"
)

const gre string = "(\\d|[a-z]+)(\\*|\\+)(\\d|[a-z]+)"

func parsetree (expr string) {
}

func main() {

        str := "2+cos*2"
        rex := regexp.MustCompile(gre)
        token := rex.FindString(str)

        fmt.Println(token)
        /*if x, err := strconv.ParseFloat(token, 64) ; err == nil {
                //fmt.Printf("%.2f \n", x*math.Cos(0))
                //str = strings.Replace(str, token, "", 1)
        }*/

}
