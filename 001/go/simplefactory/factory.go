package simplefactory

import "fmt"

// abstract product
type Chair interface {
	Sit()
}

// concrete product
type ChairRed struct {
}

func (c ChairRed) Sit() {
	fmt.Printf("sit in ChairRed\n")
}

// concrete product
type ChairBlue struct {
}

func (c ChairBlue) Sit() {
	fmt.Printf("sit in ChairBlue\n")
}

// simple factory
func ChairFactory(style string) Chair {
	switch style {
	case "blue":
		return new(ChairBlue)
	case "red":
		return new(ChairRed)
	}
	return nil
}
