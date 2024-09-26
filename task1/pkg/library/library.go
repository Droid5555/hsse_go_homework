package library

import (
	"hsse_go_homework/task1/pkg/book"
	"hsse_go_homework/task1/pkg/generator"
	"hsse_go_homework/task1/pkg/storage"
	"log"
)

type Interface interface {
	Search(string) (book.Book, bool)
	Add(book.Book) bool
	SetIdGenerator(generator.IDGenerator)
	SetStorage(storage.Interface)
}

type Library struct {
	storage  storage.Interface
	nameToID map[string]string
	hashFunc generator.IDGenerator
}

// Library methods

func (lib Library) Search(name string) (book.Book, bool) {
	b, ok := lib.storage.Search(lib.nameToID[name])
	if !ok {
		log.Println("Can't find the book with name: ", name)
	}
	return b, ok
}

func (lib *Library) Add(b book.Book) bool {
	id := lib.hashFunc(b)
	ok := false
	lib.storage, ok = lib.storage.Add(id, b)
	if ok {
		if lib.nameToID == nil {
			lib.nameToID = make(map[string]string)
		}
		lib.nameToID[b.Title] = id
		return ok
	}
	log.Println("Can't add the book: ", b)
	return ok
}

func (lib *Library) SetIdGenerator(function generator.IDGenerator) {
	if function == nil {
		return
	}

	if lib.hashFunc == nil {
		lib.hashFunc = function
		return
	}

	var newLibStorage storage.Interface

	switch lib.storage.(type) {
	case storage.BookSlice:
		newLibStorage = make(storage.BookSlice, 0)
	case storage.BookMap:
		newLibStorage = make(storage.BookMap)
	}

	var newId string
	for name, prevId := range lib.nameToID {
		b, ok := lib.storage.Search(lib.nameToID[name])
		if ok {
			newId = function(b)
			newLibStorage, _ = newLibStorage.Add(newId, b)
		} else {
			log.Println("Can't find the book at the id: ", prevId, " while transferring books to new IDs")
		}
		lib.nameToID[name] = newId
	}

	lib.storage = newLibStorage
	lib.hashFunc = function
}

func (lib *Library) SetStorage(container storage.Interface) {
	if container == nil {
		return
	}

	if lib.storage == nil {
		lib.storage = container
	}

	var newLibStorage storage.Interface

	switch container.(type) {
	case storage.BookSlice:
		newLibStorage = make(storage.BookSlice, 0)
	case storage.BookMap:
		newLibStorage = make(storage.BookMap)
	}

	for name, id := range lib.nameToID {
		b, ok := lib.storage.Search(lib.nameToID[name])
		if ok {
			newLibStorage, _ = newLibStorage.Add(id, b)
		} else {
			log.Println("Can't find the book at the id: ", id, " while transferring books to the new container")
		}
	}

	// Transferring elements from the container (it may not be empty)
	switch container.(type) {
	case storage.BookSlice:
		for _, pairIdBook := range container.(storage.BookSlice) {
			b := pairIdBook.Book
			id := lib.hashFunc(b)
			lib.nameToID[b.Title] = id
			newLibStorage, _ = newLibStorage.Add(id, b)
		}
	case storage.BookMap:
		for _, b := range container.(storage.BookMap) {
			id := lib.hashFunc(b)
			lib.nameToID[b.Title] = id
			newLibStorage, _ = newLibStorage.Add(id, b)
		}
	}
	lib.storage = newLibStorage
}
