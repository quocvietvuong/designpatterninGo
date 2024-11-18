package main

import "fmt"

type State interface {
	InsertCoin(context *VendingMachine)
	SelectItem(context *VendingMachine)
	Dispense(context *VendingMachine)
}

// Waiting State
type Waiting struct {}

func (w *Waiting) InsertCoin(context *VendingMachine) {
	fmt.Println("Coin inserted")
	context.SetState(&ItemSelected{})
}

func (w* Waiting) SelectItem(context *VendingMachine) {
	fmt.Println("Please insert a coin first")
}

func (w *Waiting) Dispense(context *VendingMachine) {
	fmt.Println("You need to select an item first")
}

// Item Selected state
type ItemSelected struct {}

func (i *ItemSelected) InsertCoin(context *VendingMachine) {
	fmt.Println("Coin already inserted")
}

func (i *ItemSelected) SelectItem(context *VendingMachine) {
	fmt.Println("Item Selected")
	context.SetState(&Dispensing{})
}

func (i *ItemSelected) Dispense(context *VendingMachine) {
	fmt.Println("You need to select an item first")
}

type Dispensing struct {}

func (d *Dispensing) InsertCoin(context *VendingMachine) {
	fmt.Println("Please wait, dispensing your item...")
}

func (d* Dispensing) SelectItem(context *VendingMachine) {
	fmt.Println("Please wait, dispensing your item...")
}

func (d *Dispensing) Dispense(context *VendingMachine) {
	fmt.Println("Dispensing item...")
	context.SetState(&Waiting{}) // Transition back to Vending Machine
}

// Vending Machine context
type VendingMachine struct {
	state State
}

func (vm *VendingMachine) SetState(state State) {
	vm.state = state
}

func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin(vm)
}

func (vm *VendingMachine) SelectItem() {
	vm.state.SelectItem(vm)
}

func (vm *VendingMachine) Dispense() {
	vm.state.Dispense(vm)
}

func main() {
	vm := &VendingMachine{}

	// Start in the Waiting state
	vm.SetState(&Waiting{})

	vm.SelectItem() // should prompt to insert a coin
	vm.InsertCoin() // Insert a coin
	vm.SelectItem() // Select an item
	vm.Dispense() // Dispense the item
	vm.InsertCoin() // Try to insert another coin while dispensing
	vm.Dispense() // Dispense while already dispensing
}