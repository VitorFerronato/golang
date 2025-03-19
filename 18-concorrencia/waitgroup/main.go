package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
