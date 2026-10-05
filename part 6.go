package main

import (
	"fmt"
	"time"
)

func search(source string, result chan<- string) {
	// Имитация задержки поиска
	time.Sleep(time.Duration(len(source)) * 100 * time.Millisecond)
	result <- fmt.Sprintf("Результат из %s", source)
}

func main() {
	sources := []string{"DB1", "DB2", "API1", "API2", "Cache"}
	result := make(chan string, len(sources))

	for _, src := range sources {
		go search(src, result)
	}

	// Получаем первый успешный результат
	fmt.Println(<-result)
	// Остальные горутины завершатся самостоятельно
}