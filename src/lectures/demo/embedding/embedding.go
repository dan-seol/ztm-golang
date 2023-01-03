package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (b Shipping) String() string {
	return []string{"Ground", "Air"}[b]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

//func (t *ToxicWaste) Convey() BeltSize {
//	return Small
//}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{4000}
	automate(&mail)
	//waste := ToxicWaste{500}
	//automate(&waste) // <- without convey method implemented this fails
}
