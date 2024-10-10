package main

import (
	"fmt"
	"time"
)

// State 1.状态集合
type State int

const (
	Red State = iota
	Yellow
	Green
)

// Event 2.事件集合
type Event string

const (
	Tick  Event = "TICK"
	Reset Event = "RESET"
)

// TrafficLight 结构体代表交通灯的状态机
type TrafficLight struct {
	currentState State
}

// NewTrafficLight 创建一个交通灯状态机实例
func NewTrafficLight() *TrafficLight {
	return &TrafficLight{currentState: Red}
}

// HandleEvent 处理指定事件并更新状态
func (t *TrafficLight) HandleEvent(event Event) {
	//根据当前状态处理输入事件
	switch t.currentState {
	case Green:
		if event == Tick {
			//绿 -> 黄
			t.currentState = Yellow
		}
	case Yellow:
		if event == Tick {
			//黄 -> 红
			t.currentState = Red
		}
	case Red:
		if event == Tick {
			//红 -> 绿
			t.currentState = Green
		}
	}

	// 打印当前状态
	fmt.Printf("Current state: %v\n", t.currentState)
}

// String 返回状态的字符串表示
func (s State) String() string {
	name := []string{"Red", "Yellow", "Green"}
	return name[s]
}
func main() {
	//实例化状态机，初始为红灯
	light := NewTrafficLight()
	fmt.Println("init :", light.currentState)

	//向状态机输入事件
	time.Sleep(5 * time.Second)
	light.HandleEvent(Tick)

	time.Sleep(3 * time.Second)
	light.HandleEvent(Tick)

	time.Sleep(5 * time.Second)
	light.HandleEvent(Tick)

	time.Sleep(5 * time.Second)
	light.HandleEvent(Tick)

	time.Sleep(5 * time.Second)
	light.HandleEvent(Reset) // 不做任何事情，因为没有定义Reset的行为
}
