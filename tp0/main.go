package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

// Constantes de las rutas
const (
	ruta1 = "archivo1.in"
	ruta2 = "archivo2.in"
)

// Lee el archivo linea por linea y devuelve un arreglo con todos los numeros que contiene dicho archivo de numeros
func leerArchivo(ruta string) []int {

	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	var arreglo []int

	for s.Scan() {

		line := s.Text()

		n, _ := strconv.Atoi(line)

		arreglo = append(arreglo, n)

	}

	return arreglo

}

// Imprime el arreglo mayor linea por linea
func imprimirArreglo(arreglo []int) {
	for _, valor := range arreglo {
		fmt.Println(valor)
	}
}

func main() {

	arreglo1 := leerArchivo(ruta1)
	arreglo2 := leerArchivo(ruta2)
	arregloMayor := ejercicios.Comparar(arreglo1, arreglo2)

	if arregloMayor == -1 {
		ejercicios.Seleccion(arreglo2)
		imprimirArreglo(arreglo2)
	} else {
		ejercicios.Seleccion(arreglo1)
		imprimirArreglo(arreglo1)
	}

}
