package main

import "fmt"

func main() {
	ticket := 50
	fName := "faiz"
	lName := "last"
	sendTicket(uint(ticket), fName, lName)
}
func sendTicket(ticket uint, fName string, lName string) {
	var UseTicket = fmt.Sprintf("%v tickets for %v %v", ticket, fName, lName)
	fmt.Printf("Sending ticket")
}
