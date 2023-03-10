package main

import "fmt"

var printed bool

func SimpleEquations(a, b, c int) {
	z := a / 3
	for {
		if z == 0 {
			if !printed { // untuk ngecek apakah sudah dicetak apa blm
				fmt.Println("no solution")
				fmt.Println(1, 2, 3)
				printed = true // set flag ke true
			}
			return
		}
		xy := b / z
		x := (a-z)/2 - xy
		y := (a-z)/2 + xy
		if x == y || y == z || x == z {
			z--
			break
		}
		if x*x+y*y+z*z == c {
			if !printed {
				fmt.Println(x, y, z)
				printed = true
			}
			return
		}
		z--
	}
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
