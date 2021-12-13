package main

import "fmt"

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		if err = f(b); err != nil {
			fmt.Printf("Error encountered: %v", err)
		}
	}
}

func (l *Library) CreateIterator() Iterator {
	return &BookIterator{
		Current: 0,
		Books:   l.Collection,
	}
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolstoy",
			PageCount: 500,
			Type:      HardCover,
		},
		{
			Name:      "The Fountainhead",
			Author:    "Ayn Rand",
			PageCount: 600,
			Type:      SoftCover,
		},
		{
			Name:      "Siddhartha",
			Author:    "Hermann Hesse",
			PageCount: 300,
			Type:      EBook,
		},
		{
			Name:      "Life",
			Author:    "Mayank Gupta",
			PageCount: 1,
			Type:      PaperBack,
		},
	},
}
