package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remaingTickets uint = 50
	//arrays
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remaingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")

	//data types
	var userTickets int
	var firstName string
	var lastName string
	var email string

	//ask user input
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	//reduce tickets

	remaingTickets = remaingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaing for %v\n",remaingTickets, conferenceName)
	// fmt.Printf("User %v booked %v tickets.\n",firstName,userTickets)
	// fmt.Printf("conferenceTickets is %T, remaingTickets is %T and conferenceName is %T\n",conferenceTickets, remaingTickets, conferenceName)

	bookings[0] = firstName +" "+lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value of the array: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The size of the array: %v\n", len(bookings))
}
