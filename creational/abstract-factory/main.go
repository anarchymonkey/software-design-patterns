package main

import "fmt"

// Abstract Factory interface
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factory 1
type WindowsGUIFactory struct{}

func (w WindowsGUIFactory) CreateButton() Button {
	return WindowsButton{}
}

func (w WindowsGUIFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// Concrete Factory 2
type MacGUIFactory struct{}

func (m MacGUIFactory) CreateButton() Button {
	return MacButton{}
}

func (m MacGUIFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}

// Abstract Product interface
type Button interface {
	Paint()
}

// Concrete Product 1
type WindowsButton struct{}

func (w WindowsButton) Paint() {
	fmt.Println("Painting a Windows button")
}

// Concrete Product 2
type MacButton struct{}

func (m MacButton) Paint() {
	fmt.Println("Painting a Mac button")
}

// Abstract Product interface
type Checkbox interface {
	Paint()
}

// Concrete Product 1
type WindowsCheckbox struct{}

func (w WindowsCheckbox) Paint() {
	fmt.Println("Painting a Windows checkbox")
}

// Concrete Product 2
type MacCheckbox struct{}

func (m MacCheckbox) Paint() {
	fmt.Println("Painting a Mac checkbox")
}

// Client code
func main() {
	windowsFactory := WindowsGUIFactory{}
	macFactory := MacGUIFactory{}

	windowsButton := windowsFactory.CreateButton()
	windowsCheckbox := windowsFactory.CreateCheckbox()

	macButton := macFactory.CreateButton()
	macCheckbox := macFactory.CreateCheckbox()

	windowsButton.Paint()
	windowsCheckbox.Paint()
	macButton.Paint()
	macCheckbox.Paint()
}
