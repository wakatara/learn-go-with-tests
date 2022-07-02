package main

type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {
	def, ok := dict[word]

	if !ok {
		return "", ErrNotFound
	}
	return def, nil

}

func (dict Dictionary) Add(word string, defn string) error {
	_, ok := dict[word]
	if !ok {
		dict[word] = defn
		return nil
	}
	return ErrWordExists
}

func (dict Dictionary) Update(word string, defn string) error {
	_, ok := dict[word]
	if ok {
		dict[word] = defn
		return nil
	}
	return ErrWordDoesNotExist
}
