package main

/*
#include <stdio.h>

void printHelloWorld() {
    puts("Hello, World from C!");
}
*/
import "C"

func main() {
	C.printHelloWorld()
}
