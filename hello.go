package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err := soma(9, 1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)

	// o operador _ (underline) pode ser usado para ignorar
	// completamento um retorno de erro porém não é uma boa prática
	// pois o idel é tratar o erro
	res, _ = soma(10, 20)
	fmt.Println(res)
}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("total maior que 10")
	}
	return res, nil
}
