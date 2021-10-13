package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name string
	Surname string
	Phone string
}

var entries []Entry

func search(key string) *Entry {
	for i, v := range entries {
		if v.Surname == key {
			return &entries[i]
		}
	}
	return nil
}

func list() {
	for _, v := range entries {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	entries = append(entries, Entry{"Mihalis", "Doe", "210941672"})
	entries = append(entries, Entry{"Mary", "Doe", "210941672"})
	entries = append(entries, Entry{"Jack", "Black", "210941672"})

	command := arguments[1]
	if command == "search" {
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}

		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	}

	if command == "list" {
		list()
	}
}