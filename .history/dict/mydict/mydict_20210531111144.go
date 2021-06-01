package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")
var errNotUpdate = errors.New("That word not found then not update")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errNotUpdate
	} else if err == nil {
		d[word] = def
	}
	return nil
}
