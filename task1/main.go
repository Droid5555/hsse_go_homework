package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	title  string
	author string
}

type PairIdBook struct {
	id   string
	book Book
}

type idGenerator func(Book) string

type BookSlice []PairIdBook
type BookMap map[string]Book

type StorageInterface interface {
	Search(string) (Book, bool)
	Add(string, Book) (StorageInterface, bool)
}

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

//Storage functions

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
	return append(bookSlice, PairIdBook{id, book}), true
}

func (bookMap BookMap) Add(id string, book Book) (StorageInterface, bool) {
	bookMap[id] = book
	return bookMap, true
}

// Library functions

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
		lib.nameToId[book.title] = id
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
			lib.nameToId[book.title] = id
			newLibStorage, _ = newLibStorage.Add(id, book)
		}
	case BookMap:
		for _, book := range container.(BookMap) {
			id := lib.hashFunc(book)
			lib.nameToId[book.title] = id
			newLibStorage, _ = newLibStorage.Add(id, book)
		}
	}
	lib.storage = newLibStorage
}

func Hash1(b Book) string {
	data := b.title + b.author
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:]) + "HASH1"
}

func Hash2(b Book) string {
	data := b.title + b.author
	hash := sha256.Sum256([]byte(data))
	hashInt := int(binary.BigEndian.Uint64(hash[:8]))
	return strconv.Itoa(hashInt) + "HASH2"
}

func Test1(lib LibraryInterface) {
	fmt.Println("\nTEST 1")
	// Создать слайс книг и библиотеку
	storage := []Book{{"Анна Каренина", "Лев Толстой"}, {"1984", "George Orwell"}}
	// загрузить книги в библиотеку
	for _, book := range storage {
		lib.Add(book)
	}
	// найти 1-2 книги в библиотеке.
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))
}

func Test2(lib LibraryInterface) {
	fmt.Println("\nTEST 2")
	// Заменить функцию генератор id
	lib.SetIdGenerator(Hash2)
	// найти еще книгу в библиотеке
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))
	lib.Add(Book{"Поющие в терновнике", "Колин Маккалоу"})
	fmt.Println(lib.Search("ПАЮЩиЕ В ТЕРНОВНИКЕ"))
	fmt.Println(lib.Search("Поющие в терновнике"))
}

func Test3(lib LibraryInterface) {
	fmt.Println("\nTEST 3")
	// Заменить хранилище
	lib.SetStorage(BookSlice{})
	// Заполнить библиотеку
	lib.Add(Book{"Детство. Отрочество. Юность.", "Лев Толстой"})
	lib.Add(Book{"Гарри Поттер и Орден Феникса", "Джоан Роулинг"})
	lib.Add(Book{"Математический анализ, интегралы и ряды", "Тюленев Александр"})
	lib.Add(Book{"How to be the best actor in the whole wide world", "Ryan Gosling"})

	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("Гарри Поттер и Узник Азкабана"))
	fmt.Println(lib.Search("Гарри Поттер и Орден Феникса"))
	fmt.Println(lib.Search("Детство. Отрочество. Юность."))
	fmt.Println(lib.Search("Ryan Gosling"))
}

func Test4(lib LibraryInterface) {
	fmt.Println("\nTEST 4 (non-empty container)")

	lib.SetStorage(BookMap{
		"1": Book{"Язык программирования С++", "Бьёрн Страуструп"},
		"2": Book{"Правила и основы игры го", "Джон Фейрбейрн"},
	})

	fmt.Println(lib.Search("Язык программирования С++"))
	fmt.Println(lib.Search("Правила и основы игры го"))
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("Детство. Отрочество. Юность."))
}

func main() {
	// lib init
	var lib LibraryInterface = &Library{}
	lib.SetIdGenerator(Hash1)
	lib.SetStorage(BookMap{})

	Test1(lib)
	Test2(lib)
	Test3(lib)
	Test4(lib)

}
