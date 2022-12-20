package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func corredor(raia chan int, name string) {
	fmt.Println("Corredor da ", name, "esta pronto para receber o bastao")
	bastao := <-raia
	fmt.Println("Corredor ", bastao, "da ", name, "esta correndo")
	if bastao == 4 {
		wg.Done()
		return
	}
	go corredor(raia, name)

	bastao++

	raia <- bastao
}

func main() {
	wg.Add(4)

	raia := make(chan int)
	raia1 := make(chan int)
	raia2 := make(chan int)
	raia3 := make(chan int)

	go func() {
		go corredor(raia, "Equipe bola")

		raia <- 1

	}()
	go func() {

		go corredor(raia1, "Equipe triangulo")

		raia1 <- 1
	}()
	go func() {
		go corredor(raia2, "Equipe quadrado")

		raia2 <- 1

	}()
	go func() {

		go corredor(raia3, "Equipe losango")

		raia3 <- 1
	}()

	wg.Wait()

	fmt.Println("A corrida acabou!")
}
