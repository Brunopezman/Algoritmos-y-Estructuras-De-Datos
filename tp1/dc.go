package main

import (
	"bufio"
	"fmt"
	"os"
	op "tp1/operaciones"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		expresion := scanner.Text()
		resultado := op.CalcularResultado(expresion)
		fmt.Println(resultado)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "ERROR", err)
		os.Exit(1)
	}
}
