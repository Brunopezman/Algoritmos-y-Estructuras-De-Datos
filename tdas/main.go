package main

import "fmt"

func elementoDesordenado(arr []int) int {
	return _elementoDesordenado(arr, 0, len(arr)-1)
}

func _elementoDesordenado(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1 // No debería llegar aquí si el arreglo cumple con la condición dada
	}

	medio := (inicio + fin) / 2

	// Verificar si el elemento en 'medio' está fuera de lugar
	if medio > inicio && arr[medio] < arr[medio-1] {
		return arr[medio]
	}
	if medio < fin && arr[medio] > arr[medio+1] {
		return arr[medio+1]
	}

	// Caso especial: si el elemento inicial es el desordenado
	if inicio < fin && arr[inicio] > arr[inicio+1] {
		return arr[inicio]
	}

	// Caso especial: si el elemento final es el desordenado
	if fin > inicio && arr[fin] < arr[fin-1] {
		return arr[fin]
	}

	// Determinar en qué mitad buscar
	if arr[inicio] <= arr[medio] {
		return _elementoDesordenado(arr, inicio, medio-1)
	} else {
		return _elementoDesordenado(arr, medio+1, fin)
	}
}

func main() {
	arr := []int{1, 2, 3, 5, 4, 6, 7, 8, 9} // 5 está fuera de lugar
	res := elementoDesordenado(arr)
	fmt.Println(res) // Debería ser 5

	arr2 := []int{1, 3, 2, 4, 5, 6, 7, 8, 9} // 3 está fuera de lugar
	res2 := elementoDesordenado(arr2)
	fmt.Println(res2) // Debería ser 3

	arr3 := []int{2, 1, 3, 4, 5, 6, 7, 8, 9} // 2 está fuera de lugar
	res3 := elementoDesordenado(arr3)
	fmt.Println(res3) // Debería ser 2

	arr4 := []int{1, 2, 3, 4, 5, 6, 7, 9, 8} // 9 está fuera de lugar
	res4 := elementoDesordenado(arr4)
	fmt.Println(res4) // Debería ser 9

	arr5 := []int{1, 3, 4, 2, 5, 6, 7, 8, 9} // 2 está fuera de lugar
	res5 := elementoDesordenado(arr5)
	fmt.Println(res5) // Debería ser 2
}
