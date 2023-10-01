package main

import "fmt"

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}
func main() {
	casey := Passenger{"Casey", 1, false}
	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	ella.Boarded = true
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4

	casey.Boarded = true
	bill.Boarded = true

	heidi.Boarded = true
	bus := Bus{heidi}

	fmt.Println(bus.FrontSeat.Name)
}
