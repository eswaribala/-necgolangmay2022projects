package models

import "fmt"

type ICustomer interface {
	View() //abstract method
}

//implementation
func (c *Customer) View() {
	fmt.Printf("%v\n", c)
}
