package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		var hey = "hey"

		fmt.Println(hey)
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v \n", firstNames)
		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out, come back next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("\nInvalid Name please try again!")
		}
		if !isValidEmail {
			fmt.Println("email address doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("The number of tickets entered is invalid address doesn't contain @ sign")
		}
		fmt.Println("Your input data is Invalid, Try again")
		// continue

	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v Application\n", conferenceName)
	fmt.Printf("we have total of %v and %v are still left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email Address: ")
	fmt.Scan(&email)
	fmt.Println("Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. Your will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
