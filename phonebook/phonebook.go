package phonebook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

type Phonebook struct {
	filename string
	entries  map[string]string
}

// Create will take a string and create a database for entries in the phonebook
func (p *Phonebook) Create(file string) (err error) {
	fh, err := os.Create(file)
	if err != nil {
		return err
	}
	p.filename = file
	if _, err := fh.Write([]byte("{}")); err != nil {
		return err
	}
	p.entries = make(map[string]string)
	return nil
}

func (p *Phonebook) Load(file string) (err error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	p.filename = file
	if err = json.Unmarshal(contents, &p.entries); err != nil {
		return err
	}
	return nil
}

func (p *Phonebook) Save() (err error) {
	if p.filename == "" {
		return errors.New("Phone database filename doesn't exist")
	}
	fh, err := os.OpenFile(p.filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	contents, err := json.Marshal(p.entries)
	if err != nil {
		return err
	}
	if _, err := fh.Write(contents); err != nil {
		return errors.New("wow: " + err.Error())
	}
	return nil
}

func (p *Phonebook) Lookup(name string) (results map[string]string, err error) {
	results = make(map[string]string)
	if p.filename == "" {
		return nil, errors.New("Phone database filename doesn't exist")
	}

	for k, v := range p.entries {
		match, err := regexp.Match(name, []byte(k))
		if err != nil {
			return nil, err
		}
		if match {
			results[k] = v
		}
	}
	return results, nil

}

func (p *Phonebook) Add(name, number string) (err error) {
	if p.filename == "" {
		return errors.New("I dont know where to find this information")
	}
	if _, ok := p.entries[name]; ok == true {
		return errors.New("Name already exists in DB")
	}
	p.entries[name] = number
	if err := p.Save(); err != nil {
		return err
	}
	return nil
}

func (p *Phonebook) Update(name, number string) (err error) {
	if p.filename == "" {
		return errors.New("I don't know where to find this information")
	}
	if _, ok := p.entries[name]; ok != true {
		return errors.New("Name doesn't exist in list, can't update")
	}
	p.entries[name] = number
	if err := p.Save(); err != nil {
		return err
	}
	return nil
}

func (p *Phonebook) Remove(name string) (err error) {
	if p.filename == "" {
		return errors.New("I don't know where to find this information")
	}
	if _, ok := p.entries[name]; ok != true {
		return errors.New("Name doesn't exist, so I can't remove it")
	}
	delete(p.entries, name)
	if err := p.Save(); err != nil {
		return err
	}
	return nil
}

func (p *Phonebook) Reverse(number string) (name string, err error) {
	if p.filename == "" {
		return "", errors.New("I don't know where to find this information")
	}
	// "Doing linear scans over an associative array is like trying to club someone to death with a loaded Uzi." - Larry Wall
	for k, v := range p.entries {
		if v == number {
			return k, nil
		}
	}
	return "", nil
}
