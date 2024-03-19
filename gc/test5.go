package main

func createLargeArray(size int) []int {
	// 这里创建了一个大小由运行时传入参数决定的切片
	return make([]int, size)
}

func main() {
	KnownSize := make([]int, 9) //可知大小：小（未逃逸）
	_ = KnownSize
	UnknownSize := make([]int, 999999) //可知大小：大（逃逸）
	_ = UnknownSize

	bigArray := createLargeArray(8) //不可知大小：无论大小（逃逸）
	_ = bigArray
}
