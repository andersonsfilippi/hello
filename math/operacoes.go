package math

var A string = "SHOWWWW"

// isso é uma função privada, ou seja, não exportada
func soma(a int, b int) int {
	return a + b
}

// Função exportada (pública)
func Soma(a int, b int) int {
	return soma(a, b)
}
