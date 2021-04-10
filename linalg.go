package mu

import (
        //"math/rand"
        "bufio"
        "encoding/csv"
        //"reflect"
)

type Vector interface {
        Magnitude() float64
}

type Matrix interface {
        Dim() (m, n)
}

func MatFromCsv(fpath string) mat Matrix {
}

func Decompose(mat *Matrix) vec Vector {
}

