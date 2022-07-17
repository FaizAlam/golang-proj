package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var city string
		// ask user for their name

		fmt.Println("Enter your first name :")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name :")
		fmt.Scan(&lastName)

		fmt.Println("Enter your city :")
		fmt.Scan(&city)
		fmt.Println("Enter your email id :")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets :")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTick := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTick {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are :%v\n", firstNames)

			if remainingTickets <= 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Println("your input data is invalid")
		}

	}

}
