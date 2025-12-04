package main

import "fmt"

// chan<- int  indica que el canal solo se usa para enviar datos(escritura)
// <-chan bool indica que el canal solo se usa para recibir datos(lectura)
func fibonacci(c chan<- int, quit <-chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // Intenta enviar x, esepaña hasta que alguien reciba
			x, y = y, x+y // Generates the sequence
		case <-quit: // Intenta recibir de quit, espera hasta que alguien envíe
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)     //recepción un canal de enteros
	quit := make(chan bool) // envío un canal de booleanos
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c) // bloquea el main hasta recibir un valor del canal c
		}
		quit <- true

	}()
	fibonacci(c, quit) // (bloquea hasta que termine la función)
}
