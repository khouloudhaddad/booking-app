package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remaingTickets uint = 50
	//arrays
	var bookings [50]string

	var sliceBooking []string

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remaingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")

	//remaingTickets >0 && len(bookings) < 50
	for  { //infinite loop

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

		//input validation
		isValidName := len(firstName) >=2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remaingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remaingTickets, userTickets)
			break
		}else{
			if !isValidName {
				fmt.Println("Name is at least two caracters !!")
			}

			if !isValidEmail{
				fmt.Println("Email address you entered doesn't contain the @ sign")
			}
		}

		//reduce tickets

		remaingTickets = remaingTickets - uint(userTickets)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaing for %v\n",remaingTickets, conferenceName)
		// fmt.Printf("User %v booked %v tickets.\n",firstName,userTickets)
		// fmt.Printf("conferenceTickets is %T, remaingTickets is %T and conferenceName is %T\n",conferenceTickets, remaingTickets, conferenceName)

		bookings[0] = firstName +" "+lastName

		sliceBooking = append(sliceBooking,firstName +" "+lastName)

		firstNames := []string{}

		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames,names[0])
		}

		fmt.Printf("The first names : %v\n",firstNames)

		if remaingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

	//data types
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value of the array: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The size of the array: %v\n", len(bookings))

	// city
	city := "London"

	switch city {
		case "New York": 
		 //execute
	case "Singapore":
		//
	case "London","Berlin":
		//
	default:
		fmt.Print("No valid city selected")	

	}
}
