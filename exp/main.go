package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}

type User struct {
	Name  string
	Dog   Dog
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
		Dog: Dog{
			Name: "Morty",
		},
		Int:   32,
		Float: 3.14,
		Slice: []string{"a", "b", "c"},
		Map: map[string]string{
			"alpha": "beta",
			"theta": "gamma",
		},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
