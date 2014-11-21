package phonebook

import (
	"os"
	"testing"
)

const dbname = "test.pb"

func TestCreatePhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Can't make phonebook: %v", err)
		return
	}

	if err := os.Remove(dbname); err != nil {
		t.Errorf("Can't remove phonebook: %v", err)
		return
	}
}

func TestLoadPhonebook(t *testing.T) {
	var p Phonebook
	// make a fake phonebook to load
	fh, err := os.Create(dbname)
	if err != nil {
		t.Errorf("Couldn't create temp db: %v", err)
		return
	}
	if _, err := fh.Write([]byte(`{"Stan Schwertly":"609-385-7359"}`)); err != nil {
		t.Errorf("Couldn't write contents to temp db: %v", err)
		return
	}
	// load phonebook
	if err := p.Load(dbname); err != nil {
		t.Errorf("Couldn't load db: %v", err)
		return
	}

	if p.entries["Stan Schwertly"] != "609-385-7359" {
		t.Error("Couldn't find me in the file, bad logical load")
		t.Errorf("%v", p.entries["Stan Schwertly"])
	}
}

func TestAddPhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Couldn't create phonebook: %v", err)
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err != nil {
		t.Errorf("Couldn't add me: %v", err)
		return
	}
	if err := p.Load(dbname); err != nil {
		t.Errorf("Couldn't load file after adding: %v", err)
		return
	}
	if p.entries["Stan Schwertly"] != "609-385-7359" {
		t.Error("I dont exist after adding\n")
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err == nil {
		t.Error("I should have failed on trying to add a name twice but didn't")
		return
	}
}

func TestLookupPhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Couldn't create phonebook: %v", err)
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err != nil {
		t.Errorf("Couldn't add me: %v", err)
		return
	}
	matches, err := p.Lookup("Stan")
	if err != nil {
		t.Errorf("Couldn't find me: %v", err)
		return
	}
	if len(matches) != 1 {
		t.Errorf("Matches is weird: %#v", matches)
		return
	}
	matches, err = p.Lookup("wow")
	if err != nil {
		t.Errorf("Got an error where should have gotten empty result: %v", err)
		return
	}
	if len(matches) != 0 {
		t.Errorf("Got results when should have gotten empty: %#v", matches)
		return
	}
}

func TestChangePhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Couldn't create phonebook: %v", err)
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err != nil {
		t.Errorf("Couldn't add me: %v", err)
		return
	}
	if err := p.Update("Stan Schwertly", "New stuff"); err != nil {
		t.Errorf("Couldn't update me: %v", err)
		return
	}
	if err := p.Update("doesn't exist", "wow"); err == nil {
		t.Error("Updated someone that doesn't exist")
		return
	}
}

func TestRemovePhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Couldn't create phonebook: %v", err)
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err != nil {
		t.Errorf("Couldn't add me: %v", err)
		return
	}
	if err := p.Remove("Stan Schwertly"); err != nil {
		t.Errorf("Couldn't remove me: %v", err)
		return
	}
	if err := p.Remove("doesn't exist"); err == nil {
		t.Error("Removed someone that doesn't exist")
		return
	}
	if err := p.Load(dbname); err != nil {
		t.Errorf("Couldn't load DB after removing: %v", err)
		return
	}
	if len(p.entries) != 0 {
		t.Errorf("DB should be empty, instead is: %v", len(p.entries))
		return
	}
}

func TestReversePhonebook(t *testing.T) {
	var p Phonebook
	if err := p.Create(dbname); err != nil {
		t.Errorf("Couldn't create phonebook: %v", err)
		return
	}
	if err := p.Add("Stan Schwertly", "609-385-7359"); err != nil {
		t.Errorf("Couldn't add me: %v", err)
		return
	}
	name, err := p.Reverse("609-385-7359")
	if err != nil {
		t.Errorf("Couldn't find me: %v", err)
		return
	}
	if name != "Stan Schwertly" {
		t.Errorf("Didn't find me, found: %v", name)
		return
	}
	name, err = p.Reverse("wow")
	if name != "" {
		t.Error("We found something we shouldn't")
	}
	if err != nil {
		t.Errorf("Should have just found nothing: %v", err)
	}

}
