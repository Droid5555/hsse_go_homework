package utils

import (
	"fmt"
	"hsse_go_homework/task1/structs"
)

func Test1(lib structs.LibraryInterface) {
	fmt.Println("\nTEST 1")
	// Создать слайс книг и библиотеку
	storage := []structs.Book{{"Анна Каренина", "Лев Толстой"}, {"1984", "George Orwell"}}
	// загрузить книги в библиотеку
	for _, book := range storage {
		lib.Add(book)
	}
	// найти 1-2 книги в библиотеке.
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))
}

func Test2(lib structs.LibraryInterface) {
	fmt.Println("\nTEST 2")
	// Заменить функцию генератор id
	lib.SetIdGenerator(Hash2)
	// найти еще книгу в библиотеке
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("1984"))
	lib.Add(structs.Book{Title: "Поющие в терновнике", Author: "Колин Маккалоу"})
	fmt.Println(lib.Search("ПАЮЩиЕ В ТЕРНОВНИКЕ"))
	fmt.Println(lib.Search("Поющие в терновнике"))
}

func Test3(lib structs.LibraryInterface) {
	fmt.Println("\nTEST 3")
	// Заменить хранилище
	lib.SetStorage(structs.BookSlice{})
	// Заполнить библиотеку
	lib.Add(structs.Book{Title: "Детство. Отрочество. Юность.", Author: "Лев Толстой"})
	lib.Add(structs.Book{Title: "Гарри Поттер и Орден Феникса", Author: "Джоан Роулинг"})
	lib.Add(structs.Book{Title: "Математический анализ, интегралы и ряды", Author: "Тюленев Александр"})
	lib.Add(structs.Book{Title: "How to be the best actor in the whole wide world", Author: "Ryan Gosling"})

	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("Гарри Поттер и Узник Азкабана"))
	fmt.Println(lib.Search("Гарри Поттер и Орден Феникса"))
	fmt.Println(lib.Search("Детство. Отрочество. Юность."))
	fmt.Println(lib.Search("Ryan Gosling"))
}

func Test4(lib structs.LibraryInterface) {
	fmt.Println("\nTEST 4 (non-empty container)")

	lib.SetStorage(structs.BookMap{
		"1": structs.Book{Title: "Язык программирования С++", Author: "Бьёрн Страуструп"},
		"2": structs.Book{Title: "Правила и основы игры го", Author: "Джон Фейрбейрн"},
	})

	fmt.Println(lib.Search("Язык программирования С++"))
	fmt.Println(lib.Search("Правила и основы игры го"))
	fmt.Println(lib.Search("Анна Каренина"))
	fmt.Println(lib.Search("Детство. Отрочество. Юность."))
}
