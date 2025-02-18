package main

import "fmt"

// Coffee interface
type Coffee interface {
	Cost() float64
	Description() string
}

// BasicCoffee struct
type BasicCoffee struct{}

func (b *BasicCoffee) Cost() float64 {
	return 2.0 // Base cost of coffee
}

func (b *BasicCoffee) Description() string {
	return "Basic Coffee"
}

// CoffeeDecorator struct
type CoffeeDecorator struct {
	Coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.Coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.Coffee.Description()
}

// MilkDecorator struct
type MilkDecorator struct {
	CoffeeDecorator
}

func (m *MilkDecorator) Cost() float64 {
	return m.CoffeeDecorator.Cost() + 0.5 // Add cost of milk
}

func (m *MilkDecorator) Description() string {
	return m.CoffeeDecorator.Description() + ", with Milk"
}

// SugarDecorator struct
type SugarDecorator struct {
	CoffeeDecorator
}

func (s *SugarDecorator) Cost() float64 {
	return s.CoffeeDecorator.Cost() + 0.2 // Add cost of sugar
}

func (s *SugarDecorator) Description() string {
	return s.CoffeeDecorator.Description() + ", with Sugar"
}

func main() {
	// Create a basic coffee
	myCoffee := &BasicCoffee{}
	fmt.Println(myCoffee.Description(), "costs $", myCoffee.Cost())

	// Decorate with milk
	myCoffeeWithMilk := &MilkDecorator{CoffeeDecorator{myCoffee}}
	fmt.Println(myCoffeeWithMilk.Description(), "costs $", myCoffeeWithMilk.Cost())

	// Decorate with sugar
	myCoffeeWithMilkAndSugar := &SugarDecorator{CoffeeDecorator{myCoffeeWithMilk}}
	fmt.Println(myCoffeeWithMilkAndSugar.Description(), "costs $", myCoffeeWithMilkAndSugar.Cost())
}
