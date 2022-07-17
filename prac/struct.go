package main

import (
	"fmt"
)

type UserData struct {
	fName   string
	lName   string
	email   string
	tickets int
}

func main() {

	//var booking = make([]UserData, 0)

	firstName := "Faiz"
	lastName := "Alam"
	email := "fa@al.com"
	ticket := 555

	var userdata = UserData{
		fName:   firstName,
		lName:   lastName,
		email:   email,
		tickets: ticket,
	}

	fmt.Println(userdata)
	fmt.Printf("%T\n", userdata)
	fmt.Println(userdata.email)
	fmt.Printf("%T\n", userdata.email)
	fmt.Println(userdata.fName)
	fmt.Printf("%T\n", userdata.fName)
	fmt.Println(userdata.tickets)
	fmt.Printf("%T\n", userdata.tickets)

}
