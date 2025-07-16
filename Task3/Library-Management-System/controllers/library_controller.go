package controllers

import (
	"fmt"
	"library-Management-System/models"
	"library-Management-System/services"
)

func Start(choice int , library *services.Library) bool {
	fmt.Println()
	switch choice {
	case 1:
		fmt.Println("Adding a new book:")
		var title, author string
		var id int
		fmt.Print("Enter Book ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter Book Title: ")
		fmt.Scan(&title)
		fmt.Print("Enter Book Author: ")
		fmt.Scan(&author)
		library.AddBook(models.Book{
			ID:     id,
			Title:  title,
			Author: author,
			Status: "available",
		})
		return true
	case 2:
		fmt.Println("Adding a new member:")
		var name string
		var id int
		fmt.Print("Enter Member ID: ")
		fmt.Scan(&id)
		fmt.Print("Enter Member Name: ")
		fmt.Scan(&name)
		library.AddMember(models.Member{
			ID:   id,
			Name: name,
		})
		return true
	case 3:
		fmt.Println("Removing a book:")
		var bookID , memberID int
		fmt.Print("Enter Book ID to remove: ")
		fmt.Scan(&bookID)
		fmt.Print("Enter Member ID: ")
		fmt.Scan(&memberID)
		library.RemoveBook(memberID , bookID)
		return true
	case 4:
		var memberID, bookID int
		fmt.Print("Enter Member ID: ")
		fmt.Scan(&memberID)
		fmt.Print("Enter Book ID: ")
		fmt.Scan(&bookID)
		library.BorrowBook(memberID, bookID)
		return true
	case 5:
		var memberID, bookID int
		fmt.Print("Enter Member ID: ")
		fmt.Scan(&memberID)
		fmt.Print("Enter Book ID: ")
		fmt.Scan(&bookID)
		library.ReturnBook(memberID, bookID)
		return true
	case 6:
		library.ListAvailableBooks()
		return true
	case 7:
		library.ListBorrowedBooks()
		return true
	case 8:
		library.ListMembers()
		return true	
	case 9:
		fmt.Println("Exiting the Library Management System. Goodbye!")
		return false
	default:
		fmt.Println("Invalid choice, please try again.")
		return true
	}
}