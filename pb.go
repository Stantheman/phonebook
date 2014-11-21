package main

import (
	"flag"
	"fmt"
	"github.com/Stantheman/phonebook/phonebook"
)

func main() {
	var p phonebook.Phonebook
	flag.Parse()
	commands := flag.Args()
	if len(commands) < 2 {
		return
	}

	// create hsphonebook.pb
	if commands[0] == "create" {
		if err := p.Create(commands[1]); err != nil {
			fmt.Printf("Couldn't create phonebook: %v\n", err)
			return
		}
		fmt.Printf("Phonebook %v created\n", commands[1])
		return
	}

	// everything from here on out needs a load, do it now
	if err := p.Load(commands[len(commands)-1]); err != nil {
		fmt.Printf("Couldn't load phonebook: %v\n", err)
		return
	}

	// lookup Sarah hsphonebook.pb
	if commands[0] == "lookup" {
		results, err := p.Lookup(commands[1])
		if err != nil {
			fmt.Printf("Couldnt lookup name: %v\n", err)
		}
		for k, v := range results {
			fmt.Printf("%v: %v\n", k, v)
		}
		return
	}

	// add 'John Michael' '123 456 4323' hsphonebook.pb
	if commands[0] == "add" {
		if err := p.Add(commands[1], commands[2]); err != nil {
			fmt.Printf("Couldn't add name: %v\n", err)
			return
		}
	}
	// change 'John Michael' '234 521 2332' hsphonebook.pb
	if commands[0] == "change" {
		if err := p.Update(commands[1], commands[2]); err != nil {
			fmt.Printf("Couldn't update name: %v\n", err)
			return
		}
	}
	// remove 'John Michael' hsphonebook.pb
	if commands[0] == "remove" {
		if err := p.Remove(commands[1]); err != nil {
			fmt.Printf("Couldn't remove name: %v\n", err)
			return
		}
	}
	// reverse-lookup '312 432 5432' hsphonebook.pb
	if commands[0] == "reverse-lookup" {
		name, err := p.Reverse(commands[1])
		if err != nil {
			fmt.Printf("Couldn't look up number: %v", err)
			return
		}
		if name == "" {
			fmt.Println("Couldn't find this number")
			return
		}
		fmt.Printf("%v matches %v\n", commands[1], name)
		return
	}
}
