package main

import "fmt"

type Passanger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passanger
}

func main() {
	casey := Passanger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passanger{Name: "Bill", TicketNumber: 2}
		ella = Passanger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passanger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has borded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
