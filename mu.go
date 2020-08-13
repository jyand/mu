package main
import (
        "fmt"
       "math"
        "regexp"
//        "strconv"
        "strings"
)

/* Global Regular Expressions */
/* used to maintain order of operations and distinguish numerical values from operators */
const GRE string = "(\\d|[a-z]+)(\\^|\\*|\\/|\\+|\\-)(\\d|[a-z]+)"
const GRE_POW string = "(\\d|[a-z]+)(\\^)(\\d|[a-z]+)"
const GRE_MULTDIV string = "(\\d|[a-z]+)(\\*|\\/)(\\d|[a-z]+)"
const GRE_ADDSUB string = "(\\d|[a-z]+)(\\+|\\-)(\\d|[a-z]+)"
const GRE_NUM string = "\\d"
const GRE_ARITH string = "\\^|\\*|\\/|\\+|\\-"

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

func subtokenizer (substr string) float64 {
        num := regexp.MustCompile(GRE_NUM)
        arith := regexp.MustCompile(GRE_ARITH)
        a := num.FindString(substr)
        substr := strings.Replace(substr, a, "")
        op := arith.FindString(substr)
        substr := strings.Replace(substr, op, "")
        b := num.FindString(substr)
        return op_eval(a, b, op)
}

func parsetree (expr string) string {
        pow := regexp.MustCompile(GRE_POW)
        multdiv := regexp.MustCompile(GRE_MULTDIV)
        addsub := regexp.MustCompile(GRE_ADDSUB)
        if pow.MatchString(expr){
                subtoken := subtokenizer(pow.FindString(GRE_POW))
                expr := strings.Replace(expr, subtoken)
                return parsetree(expr)
        }
        if multdiv.MatchString(expr){
                subtoken := subtokenizer(multdiv.FindString(GRE_MULTDIV))
                expr := strings.Replace(expr, subtoken)
                return parsetree(expr)
        }
        if addsub.MatchString(expr){
                subtoken := subtokenizer(addsub.FindString(GRE_ADDSUB))
                expr := strings.Replace(expr, subtoken)
                return parsetree(expr)
        }
        else {
                return expr
        }
}

func main() {

        str := "2+3*2"
        //str := "2+cos*2"
        rex := regexp.MustCompile(GRE)
        token := rex.FindString(str)

                /* single-variable, real-valued, differentiable functions */
/*        f := map[string]func(float64) float64 {
                "cos": math.Cos, "acos": math.Acos, "cosh": math.Cosh, "acosh": math.Acosh,
                "sin": math.Cos, "asin": math.Acos, "sinh": math.Cosh, "asinh": math.Acosh,
                "tan": math.Cos, "atan": math.Acos, "tanh": math.Cosh, "atanh": math.Acosh,
                "e^": math.Exp, "log": math.Log2, "ln": math.Log,
        }
*/
        fmt.Println(token + subtoken)
        /*if x, err := strconv.ParseFloat(token, 64) ; err == nil {
                //fmt.Printf("%.2f \n", x*math.Cos(0))
                //str = strings.Replace(str, token, "", 1)
        }*/

}
