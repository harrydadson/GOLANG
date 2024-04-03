package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const confTickets int = 50
var confName = "Go Conference"
var remainTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main()  {

	greetUsers()

	//for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber :=  helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainTickets)
	
	if isValidName && isValidEmail && isValidTicketNumber{
		
		bookTicket(uint(userTickets), firstName, lastName, email, confName)

		wg.Add(1)
		go sendTicket(uint(userTickets), firstName, lastName, email) // goroutine - concurrency

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)
		
		noTicketsRemain := remainTickets == 0
		if noTicketsRemain {
			// end program
			fmt.Println("Our conf is booked out Come back next year.")
			//break
		}
	} else if userTickets == int(remainTickets) {
		// do something els
	} else {
		if !isValidName{
			fmt.Println("first name or last name is invalid")
		} 
		if !isValidEmail{
			fmt.Println("email is invalid")
		}
		if !isValidTicketNumber{
			fmt.Println("ticket number is invalid")
		}
	}
	wg.Wait()
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first slice: %v\n", bookings[0])
	// fmt.Printf("slice type: %T\n", bookings)
	// fmt.Printf("slice length: %v\n", len(bookings))
}
//}

func greetUsers() {
	fmt.Println("Welcome to", confName, "booking application.")
	fmt.Println("We have total of", confTickets, " tickets and", remainTickets,"are still available")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ in place of index
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

		// ask for user name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your wmail address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, confName string) {
	remainTickets = remainTickets - uint(userTickets)
	
	// create a map for a user
	var userData = UserData {
		firstName: firstName, 
		lastName: lastName, 
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainTickets, confName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################################################")
}

