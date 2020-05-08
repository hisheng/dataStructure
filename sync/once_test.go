package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
once 只执行一次
*/
func TestOnce(t *testing.T) {
	var once sync.Once
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(log)
		}()
		fmt.Println("di")
	}

	time.Sleep(time.Hour)
}

func log() {
	fmt.Println("hi")
}
