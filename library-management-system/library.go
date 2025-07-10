package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	libraryInstance *Library
	once            sync.Once
)

type Library struct {
	Books   map[int]*Book
	Members map[int]*Member
}

func GetLibraryInstance() *Library {
	once.Do(func() {
		libraryInstance = &Library{Books: make(map[int]*Book), Members: make(map[int]*Member)}
	})
	return libraryInstance
}

func (l *Library) AddBook(book *Book) {
	l.Books[book.BookId] = book
	fmt.Printf("book %s has been added\n", book.Title)
}

func (l *Library) AddMember(member *Member) {
	l.Members[member.MemberId] = member
	fmt.Printf("member %s has been added\n", member.Name)
}

func (l *Library) DisplayAvailableBooks() {
	fmt.Println("Available Books:")
	for _, book := range l.Books {
		if book.IsBookAvailable() {
			fmt.Printf("book id: %d, book titile: %s\n", book.BookId, book.Title)
		}
	}
}

func (l *Library) BorrowBookByMember(bookId int, memberId int) (*BookItem, error) {
	if l.Members[memberId] == nil || l.Books[bookId] == nil {
		fmt.Println("member or book not found")
	}

	book := l.Books[bookId]
	if !book.IsBookAvailable() {
		return nil, errors.New("book not available")
	}

	member := l.Members[memberId]
	if member.IsQuotaFill() {
		return nil, errors.New("member quota fill")
	}

	borrowBook := book.BorrowBook()
	member.AddBorrowBook(borrowBook)
	fmt.Printf("member %d has been borrowed book %d\n", memberId, bookId)
	return borrowBook, nil
}

func (l *Library) ReturnBookByMember(memberId int, bookItemId int) {
	member := l.Members[memberId]

	for _, book := range member.CurrentBorrowed {
		if book.ID == bookItemId {
			book.ReturnBook()
			
		}
	}
}
