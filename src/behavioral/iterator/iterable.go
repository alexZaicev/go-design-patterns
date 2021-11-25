package main

type IterableCollection interface {
	createIterator() Iterator
}

type Iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	current int
	books []Book
}

func (bIter *BookIterator) hasNext() bool {
	return bIter.current < len(bIter.books)
}

func (bIter *BookIterator) next() *Book {
	if bIter.hasNext() {
		book := bIter.books[bIter.current]
		bIter.current += 1
		return &book
	}
	return nil
}
