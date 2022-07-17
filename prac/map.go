package main

import (
	"fmt"
)

// key : value

func main() {

	// var userData = make(map[string]string) //single map
	var userData1 = make([]map[string]string, 0) //list of maps
	for i := 0; i < 4; i++ {
		var intMap = make(map[string]string)

		var name string
		var Address string
		var number string
		fmt.Print("name : ")
		fmt.Scan(&name)
		//fmt.Println(name)
		fmt.Print("Address : ")
		fmt.Scan(&Address)
		fmt.Print("Ph no. : ")
		fmt.Scan(&number)

		intMap["name"] = name
		intMap["Address"] = Address
		intMap["Ph no."] = number

		userData1 = append(userData1, intMap)

	}
	fmt.Print(userData1)
}

/*
var firstName string
	var lastName string
	var email string
	var tickets int

	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email id :")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets :")
	fmt.Scan(&tickets)

	var userData = make(map[string]string) //single map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticket"] = strconv.FormatUint(uint64(tickets), 10)

	fmt.Print(userData)

*/
