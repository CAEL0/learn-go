package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

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
		return nil
	}
	return errWordExists
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantUpdate
	}
	d[word] = def
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return err
	}
	delete(d, word)
	return nil
}