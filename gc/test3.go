package main

var x3 int

func fun3() *int {
	x3 = 3
	return &x3
}

func main() {
	resX := fun3()
	_ = resX
}
