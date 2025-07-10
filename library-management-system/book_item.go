package main

type Status string

const (
	Available Status = "Available"
	Borrowed  Status = "Borrowed"
)

type BookItem struct {
	ID     int
	BookID int
	Status Status
}

func NewBookItem(id int, bookId int) *BookItem {
	return &BookItem{ID: id, BookID: bookId, Status: Available}
}
func (bi *BookItem) BorrowBook() {
	bi.Status = Borrowed
}

func (bi *BookItem) ReturnBook() {
	bi.Status = Available
}
