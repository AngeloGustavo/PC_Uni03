package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func corredor(raia chan int, name string, vencedor *string) {
	bastao := <-raia

	if bastao == 1 {
		fmt.Println("Equipe", name, "está preparada")
	} else {
		fmt.Println("Corredor", bastao, "da equipe", name, "está pronto para receber o bastão")
	}

	fmt.Println("Corredor", bastao, "da equipe", name, "esta correndo")
	if bastao == 4 {
		fmt.Println("\\o/ Equipe", name, "terminou a corrida!")
		if *vencedor == "" {
			*vencedor = name
		}
		wg.Done()
		return
	}
	go corredor(raia, name, vencedor)

	bastao++

	raia <- bastao
}

func main() {
	vencedor := ""
	wg.Add(4)

	raia := make(chan int)
	raia1 := make(chan int)
	raia2 := make(chan int)
	raia3 := make(chan int)

	go func() {
		go corredor(raia, "Bola", &vencedor)

		raia <- 1

	}()
	go func() {

		go corredor(raia1, "Triângulo", &vencedor)

		raia1 <- 1
	}()
	go func() {
		go corredor(raia2, "Quadrado", &vencedor)

		raia2 <- 1

	}()
	go func() {

		go corredor(raia3, "Losango", &vencedor)

		raia3 <- 1
	}()

	wg.Wait()

	fmt.Println("A corrida acabou!")
	fmt.Println()
	fmt.Println("===== Equipe", vencedor, "é a grande campeã!!! =====")
	fmt.Println()
}
