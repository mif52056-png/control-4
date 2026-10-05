package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Воркер %d обработал задачу: %s\n", id, job)
	}
}

func main() {
	numWorkers := 3
	jobs := make(chan string, 10)
	var wg sync.WaitGroup

	// Создаём воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Отправляем задачи
	for i := 1; i <= 10; i++ {
		jobs <- fmt.Sprintf("Задача %d", i)
	}
	close(jobs)

	wg.Wait()
}