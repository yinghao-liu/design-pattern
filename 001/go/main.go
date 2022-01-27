package main

import "fmt"

// abstract product
type Chair interface {
	sit()
}

// concrete product
type ChairRed struct {
}

func (c ChairRed) sit() {
	fmt.Printf("sit in ChairRed\n")
}

// concrete product
type ChairBlue struct {
}

func (c ChairBlue) sit() {
	fmt.Printf("sit in ChairBlue\n")
}

// concrete factory
func ChairFactory(style string) Chair {
	switch style {
	case "blue":
		return new(ChairBlue)
	case "red":
		return new(ChairRed)
	}

	return nil
}

// client
func main() {

	a := ChairFactory("blue")
	a.sit()
}
