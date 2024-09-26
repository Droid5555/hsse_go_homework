package library

type StorageInterface interface {
	Search(string) (Book, bool)
	Add(string, Book) (StorageInterface, bool)
}

// BookSlice Storage type 1
type BookSlice []PairIDBook

// BookMap Storage type 2
type BookMap map[string]Book

//Storage methods

func (bookSlice BookSlice) Search(id string) (Book, bool) {
	for _, pairIdBook := range bookSlice {
		if pairIdBook.id == id {
			return pairIdBook.book, true
		}
	}
	return Book{}, false
}

func (bookMap BookMap) Search(id string) (Book, bool) {
	b, ok := bookMap[id]
	return b, ok
}

func (bookSlice BookSlice) Add(id string, book Book) (StorageInterface, bool) {
	return append(bookSlice, PairIDBook{id, book}), true
}

func (bookMap BookMap) Add(id string, book Book) (StorageInterface, bool) {
	bookMap[id] = book
	return bookMap, true
}
