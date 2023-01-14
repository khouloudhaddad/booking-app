package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaingTickets = 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remaingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")

	//data types
	var userName string
	var userTickets int
	userName = "Khouloud"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n",userName,userTickets)
	fmt.Printf("conferenceTickets is %T, remaingTickets is %T and conferenceName is %T\n",conferenceTickets, remaingTickets, conferenceName)



}
