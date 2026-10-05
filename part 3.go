package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond) // 5 раз в секунду
	for i := 1; i <= 15; i++ {
		<-tick
		fmt.Printf("Запрос %d обработан в %v\n", i, time.Now().Format("15:04:05.000"))
	}
}