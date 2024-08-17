package operaciones

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	TDAPila "tdas/pila"
)

const (
	SUMA     string = "+"
	RESTA    string = "-"
	MULTI    string = "*"
	DIVISION string = "/"
	EXP      string = "^"
	RAD      string = "sqrt"
	LOG      string = "log"
	TERNARIA string = "?"
)

func CalcularResultado(expresion string) string {
	tokens := strings.Fields(expresion)
	pila := TDAPila.CrearPilaDinamica[int64]()

	for _, token := range tokens {
		switch token {
		case SUMA:
			res, err := suma(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case RESTA:
			res, err := resta(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case MULTI:
			res, err := multiplicacion(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case DIVISION:
			res, err := division(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case EXP:
			res, err := exponenciacion(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case RAD:
			res, err := radicalizacion(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case LOG:
			res, err := log(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		case TERNARIA:
			res, err := operacionTernaria(pila)
			if err != "" {
				return "ERROR"
			}
			pila.Apilar(res)
		default:
			num, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				return "ERROR"
			}
			pila.Apilar(num)
		}
	}

	if pila.EstaVacia() {
		return "ERROR"
	}
	resultado := pila.Desapilar()
	if !pila.EstaVacia() {
		return "ERROR"
	}

	return fmt.Sprintf("%d", resultado)
}

func suma(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand1 := pila.Desapilar()
	return operand1 + operand2, ""
}
func resta(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand1 := pila.Desapilar()
	return operand1 - operand2, ""
}
func multiplicacion(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand1 := pila.Desapilar()
	return operand1 * operand2, ""
}
func division(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand1 := pila.Desapilar()
	if operand2 == 0 {
		return 0, "ERROR"
	}
	return operand1 / operand2, ""
}
func exponenciacion(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand1 := pila.Desapilar()
	if operand2 < 0 {
		return 0, "ERROR"
	}
	return int64(math.Pow(float64(operand1), float64(operand2))), ""
}
func radicalizacion(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operand2 := pila.Desapilar()
	if operand2 < 0 {
		return 0, "ERROR"
	}
	return int64(math.Sqrt(float64(operand2))), ""
}
func log(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	base := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	operando := pila.Desapilar()
	if operando < 1 || base < 2 {
		return 0, "ERROR"
	}
	return int64(math.Log(float64(operando)) / math.Log(float64(base))), ""
}
func operacionTernaria(pila TDAPila.Pila[int64]) (int64, string) {
	if pila.EstaVacia() {
		return 0, "ERROR"
	}

	operand2 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}

	operand1 := pila.Desapilar()
	if pila.EstaVacia() {
		return 0, "ERROR"
	}
	condicion := pila.Desapilar()
	if condicion != 0 {
		return operand1, ""
	}
	return operand2, ""
}
