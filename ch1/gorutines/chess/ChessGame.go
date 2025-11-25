package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quit chan bool // sin buffered channel, es decir, bloquea al enviar y al recibir
//con buffer , se bloquea solo al recibir si el buffer está vacío y al enviar si el buffer está lleno

func main() {
	rand.NewSource(time.Now().UnixNano()) // Inicializa la semilla aleatoria
	move := make(chan int)
	quit = make(chan bool)

	go player("Bobby Fischer", move)
	go player("Boris Spassky", move)
	// Start the move
	move <- 1

	<-quit // Bloquea hasta que salga asignado un valor
}

func player(name string, move chan int) {
	for {

		turn := <-move // Espera a que sea su turno

		n := rand.Intn(100)
		if n <= 5 && turn >= 5 {
			fmt.Printf("Player %s was check mated and loses!", name)
			quit <- true
			return
		}

		fmt.Printf("Player %s has moved. Turn %d.\n", name, turn)
		turn++ // Incrementa el turno del siguiente jugador

		time.Sleep(1 * time.Second)
		move <- turn
	}
}
