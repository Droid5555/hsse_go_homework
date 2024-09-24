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

type BookSlice []PairIdBook
type BookMap map[string]Book

type StorageInterface interface {
	Search(id string) (Book, bool)
	Add(id string, book Book) (StorageInterface, bool)
}

type LibraryInterface interface {
	Search(name string) (Book, bool)
	Add(book Book) bool
	SetIdGenerator(func(Book) string)
	SetStorage(StorageInterface)
}

type Library struct {
	storage  StorageInterface
	nameToId map[string]string
	hashFunc func(Book) string
}

//Storage functions

func (book_slice BookSlice) Search(id string) (Book, bool) {
	for _, pair_id_book := range book_slice {
		if pair_id_book.id == id {
			return pair_id_book.book, true
		}
	}
	return Book{}, false
}

func (book_map BookMap) Search(id string) (Book, bool) {
	b, ok := book_map[id]
	return b, ok
}

func (book_slice BookSlice) Add(id string, book Book) (StorageInterface, bool) {
	return append(book_slice, PairIdBook{id, book}), true
}

func (book_map BookMap) Add(id string, book Book) (StorageInterface, bool) {
	book_map[id] = book
	return book_map, true
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
	fmt.Println("Can't add the book: ", book)
	return ok
}

func (lib *Library) SetIdGenerator(function func(Book) string) {
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
			newId = lib.hashFunc(book)
			newLibStorage, _ = newLibStorage.Add(newId, book)
		} else {
			fmt.Println("Can't find the book at the id: ", prevId, " while transferring books to new IDs")
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

	switch container.(type) {
	case BookSlice:
		container = make(BookSlice, 0)
	case BookMap:
		container = make(BookMap)
	}

	for name, id := range lib.nameToId {
		book, ok := lib.storage.Search(lib.nameToId[name])
		if ok {
			container, _ = container.Add(id, book)
		} else {
			fmt.Println("Can't find the book at the id: ", id, " while transferring books to the new container")
		}
	}

	lib.storage = container
}

func Hash1(b Book) string {
	data := b.title + b.author
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
func Hash2(b Book) string {
	data := b.title + b.author
	hash := sha256.Sum256([]byte(data))
	hashInt := int(binary.BigEndian.Uint64(hash[:8]))
	return strconv.Itoa(hashInt)
}

func main() {
	fmt.Println("TEST 1")
	// Создать слайс книг и библиотеку
	storage := []Book{{"Анна Каренина", "Лев Толстой"}, {"1984", "George Orwell"}}
	var lib LibraryInterface = &Library{}
	lib.SetIdGenerator(Hash1)
	lib.SetStorage(BookMap{})
	// загрузить книги в библиотеку
	for _, book := range storage {
		lib.Add(book)
	}
	// найти 1-2 книги в библиотеке.
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))

	fmt.Println("TEST 2")
	// Заменить функцию генератор id
	lib.SetIdGenerator(Hash2)
	// найти еще книгу в библиотеке
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))
	lib.Add(Book{"Поющие в терновнике", "Колин Маккалоу"})
	fmt.Println(lib.Search("ПАЮЩиЕ В ТЕРНОВНИКЕ"))
	fmt.Println(lib.Search("Поющие в терновнике"))

	fmt.Println("TEST 3")
	// Заменить хранилище
	lib.SetStorage(BookSlice{})
	// Заполнить библиотеку
	lib.Add(Book{"Детство. Отрочество. Юность.", "Лев Толстой"})
	lib.Add(Book{"Гарри Поттер и Орден Феникса", "Джоан Роулинг"})
	lib.Add(Book{"Математический анализ, интегралы и ряды", "Тюленев Александр"})
	lib.Add(Book{"How to be the best actor in the whole wide world", "Ryan Gosling"})
	// Найти 1-2 книги в ней.
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("Гарри Поттер и Узник Азкабана"))
	fmt.Println(lib.Search("Гарри Поттер и Орден Феникса"))
	fmt.Println(lib.Search("Детство. Отрочество. Юность."))
	fmt.Println(lib.Search("Ryan Gosling"))
}