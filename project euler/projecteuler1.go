package main

import "fmt"

func main() {
	var sa int
	var sb int
	var sc int
	var sum int

	a := 3
	b := 5
	c := 15
	e := 2

	x := 333
	y := 199
	z := 66

	sa = ((x * (a + (a * x))) / e)
	sb = ((y * (b + (b * y))) / e)
	sc = ((z * (c + (c * z))) / e)

	sum = sa + sb - sc

	fmt.Printf("Jumlah %d", sum)

}
