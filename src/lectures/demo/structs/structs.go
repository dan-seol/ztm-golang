package main

import (
	"fmt"
)

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func checkBoarded(passenger Passenger) {
	message := func(test bool, m1, m2 string) string {
		if test {
			return m1
		}
		return m2
	}(passenger.Boarded, "has boarded the bus", "has not boarded yet")
	fmt.Println(passenger.Name, message)
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)

	fmt.Println(bill, ella)
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	heidi.Boarded = true

	checkBoarded(casey)
	checkBoarded(bill)
	checkBoarded(heidi)

	bus1 := Bus{heidi}
	bus2 := Bus{casey}
	fmt.Println(bus1)
	fmt.Println(bus1.FrontSeat.Name, "is in the front seat")
	fmt.Println(bus2.FrontSeat.Name, "is in the front seat")

}
