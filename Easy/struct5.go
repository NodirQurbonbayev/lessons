package loyihalar

import (
	"fmt"
)	
type OnlineTicket struct {
	TicketNumber int
	EventName    string
	Price        float64
	Status       string
}
func (t OnlineTicket) Check() string {
	if t.Status=="Aviable" {
		return "Booked"
	} else {
		return "Cancelled"
	}
}
func (t OnlineTicket) Display() {
	fmt.Printf("Ticket Number: %d\n", t.TicketNumber)
	fmt.Printf("Event Name: %s\n", t.EventName)
	fmt.Printf("Price: $%.1f\n", t.Price)
	fmt.Printf("Status: %s\n", t.Check())
}