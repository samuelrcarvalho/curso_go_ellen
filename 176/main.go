package main

import (
	"fmt"

	"github.com/samuelrcarvalho/curso_go_ellen/176/pkg/cachorro"
)

func main() {
	fmt.Println("ola")
	saida := cachorro.Idade(7)

	fmt.Println(saida)
}
