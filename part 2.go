package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// 3 воркера
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- j * j
			}
		}()
	}

	// Отправляем 10 задач
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	// Закрываем results после завершения всех воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	// Вывод результатов
	for res := range results {
		fmt.Println(res)
	}
}