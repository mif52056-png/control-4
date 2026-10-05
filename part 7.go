package main

import (
	"fmt"
)

type Command struct {
	op    string // "inc", "get", "set"
	value int
	reply chan int
}

func manager(stateChan chan Command) {
	state := 0
	for cmd := range stateChan {
		switch cmd.op {
		case "inc":
			state += cmd.value
			cmd.reply <- state
		case "get":
			cmd.reply <- state
		case "set":
			state = cmd.value
			cmd.reply <- state
		}
	}
}

func main() {
	stateChan := make(chan Command)
	go manager(stateChan)

	reply := make(chan int)

	stateChan <- Command{op: "inc", value: 5, reply: reply}
	fmt.Println("После inc 5:", <-reply)

	stateChan <- Command{op: "get", reply: reply}
	fmt.Println("Текущее значение:", <-reply)

	stateChan <- Command{op: "set", value: 10, reply: reply}
	fmt.Println("После set 10:", <-reply)

	close(stateChan)
}