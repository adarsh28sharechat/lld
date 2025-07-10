package main

import "sync"

type Book struct {
	BookId    int
	Title     string
	BookItems []BookItem
	mu        sync.Mutex
}

func NewBook(id int, title string) *Book {
	book := &Book{BookId: id, Title: title, BookItems: make([]BookItem, 0)}
	for i := 1; i <= 10; i++ {
		book.BookItems = append(book.BookItems, *NewBookItem(i, id))
	}
	return book
}

func (b *Book) IsBookAvailable() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, item := range b.BookItems {
		if item.Status == "Available" {
			return true
		}
	}
	return false
}

func (b *Book) BorrowBook() *BookItem {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i := range b.BookItems {
		if b.BookItems[i].Status == "Available" {
			b.BookItems[i].Status = "Borrowed"
			return &b.BookItems[i]
		}
	}
	return nil
}
