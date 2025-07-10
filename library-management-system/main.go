package main

import "fmt"

func main() {
	fmt.Println("library-management-system")
	library := GetLibraryInstance()

	book1 := NewBook(1, "Maths")
	book2 := NewBook(2, "Science")
	book3 := NewBook(3, "Golang")

	library.AddBook(book1)
	library.AddBook(book2)
	library.AddBook(book3)

	library.DisplayAvailableBooks()

	member1 := NewMember(1, "Adarsh")
	member2 := NewMember(2, "Prince")

	library.AddMember(member1)
	library.AddMember(member2)

	user1borrow1, err := library.BorrowBookByMember(book1.BookId, member1.MemberId)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}

	user1borrow2, err := library.BorrowBookByMember(book2.BookId, member1.MemberId)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}

	user2borrow2, err := library.BorrowBookByMember(book2.BookId, member2.MemberId)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}

	member1.DisplayCurrentBorrowBook()
	member2.DisplayCurrentBorrowBook()

	library.DisplayAvailableBooks()

	library.ReturnBookByMember(member1.MemberId, user1borrow1.ID)
	library.ReturnBookByMember(member2.MemberId, user2borrow2.ID)
	library.ReturnBookByMember(member1.MemberId, user1borrow2.ID)

	library.DisplayAvailableBooks()

}
