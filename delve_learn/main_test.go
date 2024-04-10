package main

import "testing"

func TestMain(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("失败: 期望 3, 得到 %d", result)
	}
}
