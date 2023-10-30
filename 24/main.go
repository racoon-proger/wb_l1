package main

import (
	"log"
	"math"

	"github.com/shopspring/decimal"
)

type Point struct {
	x decimal.Decimal
	y decimal.Decimal
}

// реализуем формулу нахождения расстояния между двумя точками
// d = √(x2 - x1)2 + (y2 - y1)2
func (p *Point) GetDistance(secondPoint *Point) (float64, error) {
	numOne := secondPoint.x.Sub(p.x).Pow(decimal.NewFromInt(2))
	numSecond := secondPoint.y.Sub(p.y).Pow(decimal.NewFromInt(2))
	num, _ := numOne.Add(numSecond).Float64()
	return math.Sqrt(num), nil
}

func NewPoint(x, y decimal.Decimal) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	a := NewPoint(
		decimal.NewFromInt(1),
		decimal.NewFromInt(1),
	)
	b := NewPoint(
		decimal.NewFromInt(2),
		decimal.NewFromInt(2),
	)
	distance, err := a.GetDistance(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(distance)
}
