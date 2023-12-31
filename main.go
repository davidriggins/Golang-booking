package main

import (
	"fmt"
	"sync"
	"time"
)

// var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50

// Initialize a list of maps
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	

	// Greet the users
	greetUsers()

	
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
		bookTicket(userTickets, firstName, lastName, email)


		wg.Add(1)
		// Send the tickets. Adding go in front makes it concurrent
		go sendTicket(userTickets, firstName, lastName, email)

		// Call to get first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)


		noTicketsRemaining := remainingTickets == 0 
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked. Come back again next year")
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

	wg.Wait()

	
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	//for index, booking := range bookings {
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets;
	// For arrays
	//// bookings[0] = firstName + " " + lastName
	// For slices


	// Create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")

	wg.Done()

}