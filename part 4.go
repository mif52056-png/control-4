package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://example.com",
		"https://vk.com",
	}

	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup

	// Пул из 3 воркеров
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Ошибка при запросе к %s: %v\n", url, err)
					continue
				}
				fmt.Printf("%s -> %s\n", url, resp.Status)
				resp.Body.Close()
			}
		}()
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
	wg.Wait()
}