package main

type test struct {
	a string
	b int
	c string
}

type math struct {
	x float32
	y float64
	z int64
}

type mixed struct {
	obj1 *test
	obj2 *math
}
