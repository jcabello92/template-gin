package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("template-go iniciado correctamente!")

	for {
		fmt.Printf("[%s] template-go activo y en ejecución...\n", time.Now().Format("00:00:00"))
		time.Sleep(10 * time.Second)
	}
}