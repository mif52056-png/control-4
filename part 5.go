package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	// Укажите свои файлы
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	jobs := make(chan string, len(files))
	var wg sync.WaitGroup

	// Пул из 3 воркеров
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range jobs {
				f, err := os.Open(file)
				if err != nil {
					fmt.Printf("Ошибка открытия %s: %v\n", file, err)
					continue
				}
				h := md5.New()
				if _, err := io.Copy(h, f); err != nil {
					fmt.Printf("Ошибка чтения %s: %v\n", file, err)
					f.Close()
					continue
				}
				f.Close()
				fmt.Printf("%s: %x\n", file, h.Sum(nil))
			}
		}()
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)
	wg.Wait()
}