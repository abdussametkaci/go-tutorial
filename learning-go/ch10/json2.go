package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First, Last string
}

type Book struct {
	Title       string    `json:"book_title"`
	PageCount   int       `json:"pages,string"`
	ISBN        string    `json:"-"`
	Authors     []Name    `json:"auths,omitempty"`
	Publisher   string    `json:",omitempty"`
	PublishDate time.Time `json:"pub_date"`
}

func main() {
	books := []Book{
		{
			Title:       "Leaning Go",
			PageCount:   375,
			ISBN:        "9781784395438",
			Authors:     []Name{{"Vladimir", "Vivien"}},
			Publisher:   "Packt",
			PublishDate: time.Date(2016, time.July, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			Title:       "The Go Programming Language",
			PageCount:   380,
			ISBN:        "9780134190440",
			Authors:     []Name{{"Alan", "Donavan"}, {"Brian", "Kernighan"}},
			Publisher:   "Addison-Wesley",
			PublishDate: time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			Title:       "Introducing Go",
			PageCount:   124,
			ISBN:        "978-1491941959",
			Authors:     nil,
			Publisher:   "",
			PublishDate: time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC),
		},
	}

	file, err := os.Create("book2.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := json.NewEncoder(file)
	if err := enc.Encode(books); err != nil {
		fmt.Println(err)
	}
}
