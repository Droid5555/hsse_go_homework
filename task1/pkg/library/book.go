package library

type Book struct {
	Title  string
	Author string
}

type PairIDBook struct {
	id   string
	book Book
}
