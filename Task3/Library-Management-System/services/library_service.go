package services

import (
	"fmt"
	"library-Management-System/models"
)

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Println("Book with this ID already exists.")
		return 
	}
	l.Books[book.ID] = book
	fmt.Printf("Book '%s' by '%s' added successfully.\n", book.Title, book.Author)
}

func (l *Library) AddMember(member models.Member) {
	if _, exists := l.Members[member.ID]; exists {
		fmt.Println("Member with this ID already exists.")
		return 
	}
	l.Members[member.ID] = member
	fmt.Printf("Member '%s' added successfully.\n", member.Name)
}

func (l *Library) BorrowBook(memberID, bookID int) {
	member , memberExists := l.Members[memberID]
	book, bookExists := l.Books[bookID]

	if !memberExists {
		fmt.Println("Member not found.")
		return
	}
	if !bookExists {
		fmt.Println("Book not found.")
		return
	}
	if book.Status != "available" {
		fmt.Println("Book is not available for borrowing.")
		return
	}
	book.Status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Books[bookID] = book
	l.Members[memberID] = member	
	fmt.Printf("Member '%s' borrowed book '%s'.\n", member.Name, book.Title)
}

func (l *Library) RemoveBook(memberID, bookID int) {
	member, memberExists := l.Members[memberID]
	book, bookExists := l.Books[bookID]
	if !memberExists {
		fmt.Println("Member not found.")
		return
	}
	if !bookExists {
		fmt.Println("Book not found.")
		return
	}
	
	if book.Status == "borrowed" {
		for i, b := range member.BorrowedBooks {
			if b.ID == bookID {
				member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
				break
			}
		}
	}
	delete(l.Books, bookID)
	l.Members[memberID] = member
	fmt.Printf("Member '%s' removed book '%s'.\n", member.Name, book.Title)
}

func (l *Library) ReturnBook(memberID, bookID int) {
	member, memberExists := l.Members[memberID]
	book, bookExists := l.Books[bookID]

	if !memberExists {
		fmt.Println("Member not found.")
		return
	}
	if !bookExists {
		fmt.Println("Book not found.")
		return
	}
	if book.Status != "borrowed" {
		fmt.Println("Book is not currently borrowed.")
		return
	}

	book.Status = "available"
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.Books[bookID] = book
	l.Members[memberID] = member	
	fmt.Printf("Member '%s' returned book '%s'.\n", member.Name, book.Title)
}

func displayBook(book models.Book){
	fmt.Println()
	fmt.Println("Book Details:")
	fmt.Println("----------------------------- /\\ ---------------------")
	fmt.Printf("ID: %d, Title: '%s', Author: '%s', Status: '%s'\n", book.ID, book.Title, book.Author, book.Status)
	fmt.Println("----------------------------- /\\ ---------------------")
}


func (l *Library) ListAvailableBooks() {
	fmt.Println("Available Books:")
	for _, book := range l.Books {
		if book.Status == "available" {
			displayBook(book)
		}
	}
}
func (l *Library) ListMembers() {
	fmt.Println("Members:")
	for _, member := range l.Members {
		fmt.Printf("ID: %d, Name: '%s', Borrowed Books: ", member.ID, member.Name)
		if len(member.BorrowedBooks) == 0 {
			fmt.Println("None")
		} else {
			for _, book := range member.BorrowedBooks {
				fmt.Printf("'%s' ", book.Title)
			}
			fmt.Println()
		}
	}
}

func (l *Library) ListBorrowedBooks() {
	fmt.Println("Borrowed Books:")
	for book_id := range l.Books {
		book := l.Books[book_id]
		if book.Status == "borrowed" {
			displayBook(book)
		}
	}
}

