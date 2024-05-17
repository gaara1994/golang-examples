package cpu

import (
	"fmt"
	"math"
	"time"
)

// CalculatePi 使用莱布尼茨公式计算π/4的近似值，并返回π的近似值（乘以4后）。
// 参数n表示要计算的级数项数。
func CalculatePi(n int) float64 {
	fmt.Println("开始计算π的近似值（使用莱布尼茨公式）：")
	startTime := time.Now()
	piOverFour := 0.0 // 初始化π/4的近似值
	sign := 1.0
	for i := 0; i < n; i++ {
		piOverFour += sign / float64(2*i+1)
		sign = -sign
	}
	duration := time.Since(startTime)
	fmt.Printf("计算得到的π/4的近似值: %.15f\n", piOverFour)
	fmt.Printf("数学库提供的π值: %.15f\n", math.Pi)
	fmt.Printf("计算耗时: %s\n", duration)
	// 返回π的近似值（将π/4的近似值乘以4）
	return piOverFour * 4
}