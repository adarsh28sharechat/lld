package main

import "fmt"

type Member struct {
	MemberId        int
	Name            string
	CurrentBorrowed []*BookItem
}

func NewMember(memberId int, name string) *Member {
	return &Member{MemberId: memberId, Name: name, CurrentBorrowed: make([]*BookItem, 0)}
}

func (m *Member) IsQuotaFill() bool {
	return len(m.CurrentBorrowed) >= 3
}

func (m *Member) AddBorrowBook(borrowBook *BookItem) {
	m.CurrentBorrowed = append(m.CurrentBorrowed, borrowBook)
}

func (m *Member) DisplayCurrentBorrowBook() {
	fmt.Printf("current borrowed books of member %d\n", m.MemberId)

	for _, bi := range m.CurrentBorrowed {
		fmt.Printf(" - Book %d (Item ID: %d)\n", bi.BookID, bi.ID)
	}
}

func (m *Member) RemoveBorrowBook(bookItem *BookItem) {
	for i, bi := range m.CurrentBorrowed {
		if bi.ID == bookItem.ID {
			m.CurrentBorrowed = append(m.CurrentBorrowed[:i], m.CurrentBorrowed[i+1:]...)
		}
	}
}
