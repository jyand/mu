package mu

import (
        "math/rand"
        "time"
        "bufio"
        "os"
        "strings"
)

type Vector interface {
        Magnitude() float64
}

type Matrix interface {
        Dim() (m, n int)
}

// Type switch for a vector
func VecType(vin Vector) (vout Vector, numeric bool) {
        switch t := vin.(type) {
        case int:
                v := make([]int, len(vin))
                copy(v, vin)
                return (v, true)
        case float64:
                v := make([]float64, len(vin))
                copy(v, vin)
                return (v, true)
        case uint64:
                v := make([]uint64, len(vin))
                copy(v, vin)
                return (v, true)
        case complex128:
                v := make([]int, len(vin))
                copy(v, vin)
                return (v, true)
        default:
                return (v, false)
        }
}

// defining a matrix by importing data is generally efficient and convenient
func MatFromCsv(fpath string) (mat Matrix, numeric bool) {
        file, _ := os.Open(fpath)
        defer file.Close()
        var m Matrix
        sc := bufio.NewScanner(file)
        for sc.Scan() {
                m = append(m, (strings.Split(Sc.Text(), ",")))
        }
        // apply the type switch on entire matrix as an array of vectors
        // maybe we don't need both vector and matrix interfaces?
        return VecType(m)
}

/* generate an mxn matrix of random, nonnegative real numbers within range (a, b]
useful for 'weights', Markov Chains, etc. */
func RandMat(m int, n int, a float64, b float64) mat Matrix {
        mat := [m][n]float64{}
        for j := 0 ; j <m ; j++ {
                for i := 0 ; i < n ; i++ {
                        rand.Seed(time.Now().UnixNano())
                        mat[i][j] = b*rand.float64() + a
                }
        }
}

func Transpose(min Matrix) mout Matrix {
}
