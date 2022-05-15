package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey Jones", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{"Bill Smith", 2, false}
		ella = Passenger{"Ella Fitzgerald", 3, false}
	)

	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi Montag"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill got on the bus.")
	}
	if casey.Boarded {
		fmt.Println("Casey got on the bus.")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat.")
}
