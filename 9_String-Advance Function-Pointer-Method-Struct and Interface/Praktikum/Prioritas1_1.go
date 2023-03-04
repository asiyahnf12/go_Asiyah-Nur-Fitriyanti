package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // liters per kilometer
	return c.FuelIn / fuelEfficiency
}

func main() {
	car := Car{
		Type:   "BMW",
		FuelIn: 30, // liters
	}
	distance := car.EstimateDistance()
	fmt.Printf("The %s can travel an estimated %.2f kilometers.\n", car.Type, distance)
}
