package penquin

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 基于timer 去做惩罚和处理
func Test(t *testing.T) {
	//closeCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	dataCh := make(chan int, 10)
	tm := time.NewTimer(0 * time.Second)
	i := 1
	go func() {
		time.Sleep(30 * time.Second)
		cancel()
	}()
	go func() {
		for i := 0; i <= 100; i++ {
			dataCh <- i
		}
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("close")
			return
		case <-tm.C:
			//time.Sleep(30 * time.Second)
			if rand.Int()%3 > 0 {
				tm.Reset(2 * time.Second * time.Duration(i))
				fmt.Println("tm reset", i)
				i++
				continue
			}
			tm.Reset(0 * time.Second)
			i = 1
			fmt.Println("tm not reset")
		}
	}
	tm.Stop()
}
