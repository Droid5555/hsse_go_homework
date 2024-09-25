package structs

type Book struct {
	Title  string
	Author string
}

type PairIdBook struct {
	id   string
	book Book
}

type BookSlice []PairIdBook
type BookMap map[string]Book
