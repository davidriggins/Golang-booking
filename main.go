package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}


	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	
	for {

		// For arrays
		//// var bookings = [50]string{}
		// For slices
		//// var bookings []string

		
		// Get user input
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		

		if isValidName && isValidEmail && isValidTicketNumber {
			
			// Book the tickets
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Call to get first names
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)


			noTicketsRemaining := remainingTickets == 0 
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked. Come back again next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

		
		// city := "London"

		// switch city {
		// 	case "New York":
		// 		// Execute code for booking New York
		// 	case "Singapore":
		// 		// Execute code for booking Singapore
		// 	case "London", "Berlin":
		// 		// Execute code for booking London or Berlin
		// 	default:
		// 		// When none of above are true, execute this code
		// }
	
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	//for index, booking := range bookings {
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name. Use the memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets;
	// For arrays
	//// bookings[0] = firstName + " " + lastName
	// For slices
	bookings = append(bookings, firstName + " " + lastName)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}