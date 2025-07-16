package main

import (
	"fmt"
	"library-Management-System/controllers"
	"library-Management-System/services"
)

func main() {
	choice := true
	library := services.NewLibrary()
	for choice {
		var UserChoice int
		fmt.Println()
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. View available Books")
		fmt.Println("7. View borrowed Books")
		fmt.Println("8. View Members")
		fmt.Println("9. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&UserChoice)

		choice = controllers.Start(UserChoice, library)
	}
	fmt.Println("Thank you for using the Library Management System!")
}