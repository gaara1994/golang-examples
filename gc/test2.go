package main

func fun2() *int {
	x := 10
	return &x
}

func main() {
	resX := fun2()
	_ = resX
}
