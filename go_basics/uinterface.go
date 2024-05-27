package gobasics

import "fmt"

type MotorVechical interface {
	Milegae() float64
	Range() float64
}

type BMW struct {
	fuel         float64
	distance     float64
	averagespeed int
}

type Tesla struct {
	battery      float32
	distance     float64
	averagespeed int
}

func (m BMW) Milegae() float64 {
	return m.distance / m.fuel
}

func (m BMW) Range() float64 {
	return m.Milegae() * m.fuel
}

func (t Tesla) Milegae() float64 {
	return t.distance / float64(t.battery)
}

func (t Tesla) Range() float64 {
	return t.Milegae() * float64(t.battery)
}

func TestInterface() {
	b1 := BMW{fuel: 10, distance: 20, averagespeed: 30}
	t1 := Tesla{battery: 50, distance: 20, averagespeed: 30}
	t := []MotorVechical{b1, t1}
	fmt.Println("Milage BMW:", t[0].Milegae())
	fmt.Println("Range BMW:", t[0].Range())
	fmt.Println("Milage Tesla", t[1].Milegae())
	fmt.Println("Range Tesla", t[1].Range())
}
