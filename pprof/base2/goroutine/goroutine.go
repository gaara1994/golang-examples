package goroutine

import "time"

type Goroutine struct {

}

func (g *Goroutine)Run() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Minute)
		}()
	}
}
