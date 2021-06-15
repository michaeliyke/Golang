package main

import "example.com/log"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return log.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"YearPublished: \t\t%v\n",
		b.Title, b.Author,
		b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide+ to Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tokien",
		YearPublished: 1939,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charse Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	Book{
		ID:            6,
		Title:         "I, Robot",
		Author:        "Asaac Asimov",
		YearPublished: 1950,
	},
	Book{
		ID:            7,
		Title:         "The gods Themselves",
		Author:        "Isaac Asimov",
		YearPublished: 1973,
	},
	Book{
		ID:            8,
		Title:         "The Moon is a Harsh Mistress",
		Author:        "Robert A. Heinlein",
		YearPublished: 1966,
	},
	Book{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	Book{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John scalzi",
		YearPublished: 2006,
	}}
