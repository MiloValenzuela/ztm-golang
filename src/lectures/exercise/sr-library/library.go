//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type BookID string
type MemberID string

// Book represents a liberey book
type Book struct {
	ID     BookID
	Title  string
	Author string
}

// Member represents a liberey member
type Member struct {
	ID   MemberID
	name string
}

// CheckoutRecord tracks a book's checkout status
type CheckoutRecord struct {
	BookID       BookID
	MemberID     MemberID
	CheckoutTime time.Time
	ReturnTime   time.Time // Zero value if not returned yet
}

type Library struct {
	Books     map[BookID]Book
	Members   map[MemberID]Member
	Checkouts map[BookID]CheckoutRecord
}

// Initialize the Library
func NewLibrary() *Library {
	return &Library{
		Books:     make(map[BookID]Book),
		Members:   make(map[MemberID]Member),
		Checkouts: make(map[BookID]CheckoutRecord),
	}
}

// Adding books and members
func (lib *Library) AddBook(id BookID, title, author string) {
	lib.Books[id] = Book{ID: id, Title: title, Author: author}
}

func (lib *Library) AddMember(id MemberID, name string) {
	lib.Members[id] = Member{ID: id, name: name}
}

// Check Out a Book
func (lib *Library) CheckOut(bookID BookID, memberID MemberID) error {
	// Check if book exists and isn't already checked out
	if _, exists := lib.Books[bookID]; !exists {
		return fmt.Errorf("book not found")
	}
	if _, checkOut := lib.Checkouts[bookID]; checkOut {
		return fmt.Errorf("book already checked out")
	}

	lib.Checkouts[bookID] = CheckoutRecord{
		BookID:       bookID,
		MemberID:     memberID,
		CheckoutTime: time.Now(),
	}
	return nil
}

func (lib *Library) CheckIn(bookID BookID) error {
	record, exists := lib.Checkouts[bookID]
	if !exists {
		return fmt.Errorf("book was not checked out")
	}

	record.ReturnTime = time.Now()
	lib.Checkouts[bookID] = record
	return nil
}

func (lib *Library) PrintState() {
	fmt.Println("\n--- Library State ---")
	fmt.Println("Books:")
	for _, book := range lib.Books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}

	fmt.Println("\nMembers:")
	for _, member := range lib.Members {
		fmt.Printf("- %s (ID: %s)\n", member.name, member.ID)
	}

	fmt.Println("\nChecked Out Books:")
	for _, record := range lib.Checkouts {
		book := lib.Books[record.BookID]
		member := lib.Members[record.MemberID]
		fmt.Printf("- %s (checked out by %s at %v)\n", book.Title, member.name, record.CheckoutTime)
	}
}

func main() {
	lib := NewLibrary()

	// Add books and members
	lib.AddBook("b1", "The Go Programming Language", "Alan A. A. Donovan")
	lib.AddBook("b2", "1984", "George Orwell")
	lib.AddMember("m1", "Alice")
	lib.AddMember("m2", "Bob")

	lib.PrintState()

	// Check out a book
	lib.CheckOut("b1", "m1")
	lib.PrintState()

	// Check in the book
	lib.CheckIn("b1")
	lib.PrintState()
}
