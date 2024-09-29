package book

type Book struct {
	Title  string
	Author string
}

type PairIDBook struct {
	ID   string
	Book Book
}
