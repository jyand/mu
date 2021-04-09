# μ
====================
a library with some numerical methods and some math functions not included in the main math pkg

Install
-------

go get github.com/jyand/mu

Usage
-----
mu currently has 3 different sub-packages:
* discrete: operations for natural numbers, be aware that most of the functions take uint64's and not ints 
* continuous: some misc. math functions defined on the "real" numbers
* analysis: numerical methods for calculus and ordinary differential equations
so import as "github.com/jyand/*subpackage*"

Next Up
-------
* linear algebra/matrix stuff
* root finding methods
* interpolation/extrapolation methods
* some basic optimization methods (e.g. simplex methods, calculus-based techniques)
* a few different options for regression
* fft
