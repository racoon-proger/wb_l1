package main

import (
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	a := decimal.NewFromInt(4)
	a = a.Pow(decimal.NewFromInt(20))

	b := decimal.NewFromInt(3)
	b = b.Pow(decimal.NewFromInt(20))

	// умножение
	c := a.Mul(b)
	log.Println(c)

	// деление
	c = a.Div(b)
	log.Println(c)

	// сложение
	c = a.Add(b)
	log.Println(c)

	// вычитание
	c = a.Sub(b)
	log.Println(c)
}
