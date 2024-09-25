package structs

import "log"

type LibraryInterface interface {
	Search(string) (Book, bool)
	Add(Book) bool
	SetIdGenerator(idGenerator)
	SetStorage(StorageInterface)
}

type Library struct {
	storage  StorageInterface
	nameToId map[string]string
	hashFunc idGenerator
}

// Library methods

func (lib Library) Search(name string) (Book, bool) {
	b, ok := lib.storage.Search(lib.nameToId[name])
	if !ok {
		log.Println("Can't find the book with name: ", name)
	}
	return b, ok
}

func (lib *Library) Add(book Book) bool {
	id := lib.hashFunc(book)
	ok := false
	lib.storage, ok = lib.storage.Add(id, book)
	if ok {
		if lib.nameToId == nil {
			lib.nameToId = make(map[string]string)
		}
		lib.nameToId[book.Title] = id
		return ok
	}
	log.Println("Can't add the book: ", book)
	return ok
}

func (lib *Library) SetIdGenerator(function idGenerator) {
	if function == nil {
		return
	}

	if lib.hashFunc == nil {
		lib.hashFunc = function
		return
	}

	var newLibStorage StorageInterface

	switch lib.storage.(type) {
	case BookSlice:
		newLibStorage = make(BookSlice, 0)
	case BookMap:
		newLibStorage = make(BookMap)
	}

	var newId string
	for name, prevId := range lib.nameToId {
		book, ok := lib.storage.Search(lib.nameToId[name])
		if ok {
			newId = function(book)
			newLibStorage, _ = newLibStorage.Add(newId, book)
		} else {
			log.Println("Can't find the book at the id: ", prevId, " while transferring books to new IDs")
		}
		lib.nameToId[name] = newId
	}

	lib.storage = newLibStorage
	lib.hashFunc = function
}

func (lib *Library) SetStorage(container StorageInterface) {
	if container == nil {
		return
	}

	if lib.storage == nil {
		lib.storage = container
	}

	var newLibStorage StorageInterface

	switch container.(type) {
	case BookSlice:
		newLibStorage = make(BookSlice, 0)
	case BookMap:
		newLibStorage = make(BookMap)
	}

	for name, id := range lib.nameToId {
		book, ok := lib.storage.Search(lib.nameToId[name])
		if ok {
			newLibStorage, _ = newLibStorage.Add(id, book)
		} else {
			log.Println("Can't find the book at the id: ", id, " while transferring books to the new container")
		}
	}

	// Transferring elements from the container (it may not be empty)
	switch container.(type) {
	case BookSlice:
		for _, pairIdBook := range container.(BookSlice) {
			book := pairIdBook.book
			id := lib.hashFunc(book)
			lib.nameToId[book.Title] = id
			newLibStorage, _ = newLibStorage.Add(id, book)
		}
	case BookMap:
		for _, book := range container.(BookMap) {
			id := lib.hashFunc(book)
			lib.nameToId[book.Title] = id
			newLibStorage, _ = newLibStorage.Add(id, book)
		}
	}
	lib.storage = newLibStorage
}
