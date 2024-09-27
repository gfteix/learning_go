package main

type Dictionary map[string]string

// read this https://dave.cheney.net/2016/04/07/constant-errors

// making the errors more reusable and immutable.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

// An interesting property of maps is that you can modify them without passing as an address to it (e.g

/*
So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps here.

Therefore, you should never initialize a nil map variable, instead you should initialize an empty map or use the make keyword to create a map
*/
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word) // built-in function, returns nothing
}
