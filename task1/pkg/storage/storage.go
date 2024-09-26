package storage

import "hsse_go_homework/task1/pkg/book"

type Interface interface {
	Search(string) (book.Book, bool)
	Add(string, book.Book) (Interface, bool)
}

// BookSlice Storage type 1
type BookSlice []book.PairIDBook

// BookMap Storage type 2
type BookMap map[string]book.Book

//Storage methods

func (bookSlice BookSlice) Search(id string) (book.Book, bool) {
	for _, pairIdBook := range bookSlice {
		if pairIdBook.ID == id {
			return pairIdBook.Book, true
		}
	}
	return book.Book{}, false
}

func (bookMap BookMap) Search(id string) (book.Book, bool) {
	b, ok := bookMap[id]
	return b, ok
}

func (bookSlice BookSlice) Add(id string, b book.Book) (Interface, bool) {
	return append(bookSlice, book.PairIDBook{ID: id, Book: b}), true
}

func (bookMap BookMap) Add(id string, b book.Book) (Interface, bool) {
	bookMap[id] = b
	return bookMap, true
}
