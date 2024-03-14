package main

func fun1() int {
	x := 10
	return x
}

func main() {
	var a = 5
	_ = a

	resX := fun1()
	_ = resX
}
