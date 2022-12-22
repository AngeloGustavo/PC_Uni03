package main

import (
	"fmt"
	"sync"
)

// notifica o programa para esperar por uma ação
var wg sync.WaitGroup

/*
Função usada por cada umas das 4 threads para
@param raia raia da equipe
@param nome nome da equipe
@param vencedor ponteiro que no fim da corrida recebe o nome da equipe vencedora
*/
func corredor(raia chan int, nome string, vencedor *string) {
	bastao := <-raia

	if bastao == 1 {
		fmt.Println("Equipe", nome, "está preparada")
	} else {
		fmt.Println("Corredor", bastao, "da equipe", nome, "está pronto para receber o bastão")
	}

	fmt.Println("Corredor", bastao, "da equipe", nome, "esta correndo")
	if bastao == 4 {
		fmt.Println("\\o/ Equipe", nome, "terminou a corrida!")
		if *vencedor == "" {
			*vencedor = nome
		}
		wg.Done()
		return
	}
	go corredor(raia, nome, vencedor)

	bastao++

	raia <- bastao
}

/*
Função main executável
*/
func main() {
	// string que recebe o nome da equipe vencedora
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
