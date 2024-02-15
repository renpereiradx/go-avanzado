package main

import "fmt"

type Iproduct interface {
	SetStock(stock int)
	GetStock() int
	SetName(name string)
	GetName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) SetStock(stock int) {
	c.stock += stock
}

func (c *Computer) GetStock() int {
	return c.stock
}

func (c *Computer) SetName(name string) {
	c.name = name
}

func (c *Computer) GetName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLaptop() Iproduct {
	return &Laptop{
		Computer: Computer{
			name:  "laptop",
			stock: 3,
		},
	}
}

func newDesktop() Iproduct {
	return &Laptop{
		Computer: Computer{
			name:  "desktop",
			stock: 5,
		},
	}
}

func GetComputerFactory(computerType string) (Iproduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}
	if computerType == "desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("invalid computer type")
}

func PrintData(product Iproduct) {
	fmt.Printf("Nombre: %s\n Stock: %d\n", product.GetName(), product.GetStock())
}

func main() {
	var laptop2 Laptop
	laptop2.SetName("laptop2")
	laptop2.SetStock(4)
	desktop2 := new(Desktop)
	desktop2.name = "desktop2"
	desktop2.stock = 6
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")
	PrintData(laptop)
	PrintData(desktop)
	PrintData(&laptop2)
	PrintData(desktop2)
	fmt.Println("laptop ", laptop, " laptop2 ", laptop2)
	fmt.Println("Desktop ", desktop, "desktop2 ", desktop2)
}
