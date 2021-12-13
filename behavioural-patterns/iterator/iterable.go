package main

type IterableCollection interface {
	CreateIterator() Iterator
}
type Iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	Current int
	Books   []Book
}

func (b *BookIterator) HasNext() bool {
	if b.Current < len(b.Books) {
		return true
	}
	return false
}

func (b *BookIterator) Next() *Book {
	if b.HasNext() {
		book := b.Books[b.Current]
		b.Current += 1
		return &book
	}
	return nil
}
