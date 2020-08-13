# mu
Tools for numerical computing written in Go
Useful in shell scripts

mu [SUBCOMMAND] *flags* f(x) [arg1,arg2,...]

## Motivation
* to rewrite in Go some numerical methods I've previously written in C or matlab
* a personal exercise in creating interpreters
* to investigate the potential of writing interpreters completely in Go (most are in C/C++)

## Progress
At the moment, this is a command-line program that simply evaluates mathematical functions using sort of a mini-interpreter design. I have never read a book on Compilers/Interpreters so I am approaching this very much intuitively.
The end-game is to have a much more sophisticated command-line tool that can compute derivatives, integrals, fast fourier transforms, matrix operations, etc. and make use of csv's or other formatted data as input/output.
