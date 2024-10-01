package storage

import generator "task1/generator"

type Book struct {
	ID     int
	Author string
	Name   string
}

// Интерфейс для хранения книг
type BookStorage interface {
	AddBook(name, author string)
	SearchByID(id int) *Book
	SearchByName(name string) *Book
	RefreshID(idGen *generator.IDGenerator)
	SetIDGenerator(idGen *generator.IDGenerator)
}

// Реализация хранения в слайсе
type SliceStorage struct {
	Books []Book
	idGen *generator.IDGenerator
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		Books: make([]Book, 0),
	}
}

func (s *SliceStorage) SetIDGenerator(idGen *generator.IDGenerator) {
	s.idGen = idGen
}

func (s *SliceStorage) AddBook(name, author string) {
	book := Book{
		ID:     s.idGen.GenerateID(),
		Name:   name,
		Author: author,
	}
	s.Books = append(s.Books, book)
}

func (s *SliceStorage) SearchByID(id int) *Book {
	for _, book := range s.Books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func (s *SliceStorage) SearchByName(name string) *Book {
	for _, book := range s.Books {
		if book.Name == name {
			return &book
		}
	}
	return nil
}

func (s *SliceStorage) RefreshID(idGen *generator.IDGenerator) {
	s.idGen = idGen
	for i := range s.Books {
		s.Books[i].ID = s.idGen.GenerateID()
	}
}

type MapStorage struct {
	books map[int]Book
	idGen *generator.IDGenerator
}

func (m *MapStorage) SetIDGenerator(idGen *generator.IDGenerator) {
	m.idGen = idGen
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		books: make(map[int]Book),
	}
}

func (m *MapStorage) AddBook(name, author string) {
	book := Book{
		ID:     m.idGen.GenerateID(),
		Name:   name,
		Author: author,
	}
	m.books[book.ID] = book
}

func (m *MapStorage) SearchByID(id int) *Book {
	if book, ok := m.books[id]; ok {
		return &book
	}
	return nil
}

func (m *MapStorage) SearchByName(name string) *Book {
	for _, book := range m.books {
		if book.Name == name {
			return &book
		}
	}
	return nil
}

func (m *MapStorage) RefreshID(idGen *generator.IDGenerator) {
	m.idGen = idGen
	for id, book := range m.books {
		newID := m.idGen.GenerateID()
		book.ID = newID
		delete(m.books, id)
		m.books[newID] = book
	}
}
