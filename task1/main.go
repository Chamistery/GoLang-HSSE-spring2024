package main

import (
  "fmt"
  "math/rand"
)

func generateID(customGenerate func() int) int {
  return customGenerate()
}

type Book struct {
  ID     int
  Author string
  Name   string
}

type BookStorage interface {
  AddBook(name, author string, generate func() int)
  SearchByID(id int) *Book
  SearchByName(name string) *Book
  refreshID(generate func() int)
}

type Library struct {
  storage  BookStorage
  generate func() int
}

func CreateLibrary(storage BookStorage, generate func() int) *Library {
  return &Library{
    storage:  storage,
    generate: generate,
  }
}

func (lib *Library) ChangeId(generate func() int) {
  lib.generate = generate
  lib.storage.refreshID(generate)
}

func (lib *Library) AddBook(name, author string) {
  lib.storage.AddBook(name, author, lib.generate)
}

func (lib *Library) SearchByID(id int) *Book {
  return lib.storage.SearchByID(id)
}

func (lib *Library) SearchByName(name string) *Book {
  return lib.storage.SearchByName(name)
}

type SliceStorage struct {
  books []Book
}

func NewSliceStorage() *SliceStorage {
  return &SliceStorage{books: make([]Book, 0)}
}

func (s *SliceStorage) AddBook(name, author string, generate func() int) {
  book := Book{
    ID:     generateID(generate),
    Name:   name,
    Author: author,
  }
  s.books = append(s.books, book)
}

func (s *SliceStorage) refreshID(generate func() int) {
  for i := range s.books {
    s.books[i].ID = generate()
  }
}

func (s *SliceStorage) SearchByID(id int) *Book {
  for _, book := range s.books {
    if book.ID == id {
      return &book
    }
  }
  return nil
}

func (s *SliceStorage) SearchByName(name string) *Book {
  for _, book := range s.books {
    if book.Name == name {
      return &book
    }
  }
  return nil
}

type MapStorage struct {
  books map[int]Book
}

func NewMapStorage() *MapStorage {
  return &MapStorage{books: make(map[int]Book)}
}

func (m *MapStorage) AddBook(name, author string, generate func() int) {
  book := Book{
    ID:     generateID(generate),
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

func (m *MapStorage) refreshID(generate func() int) {
  for id, book := range m.books {
    newID := generate()
    book.ID = newID
    delete(m.books, id)
    m.books[newID] = book
  }
}

var idCounter int

func generateDirectly() int {
  idCounter++
  return idCounter
}

func generateRandomly() int {
  return rand.Intn(100000)
}

func main() {
  // Slice
  sliceStorage := NewSliceStorage()
  library := CreateLibrary(sliceStorage, generateDirectly)
  library.AddBook("The Last Wish", "Andrzej Sapkowski")
  library.AddBook("Sword of Destiny", "Andrzej Sapkowski")

  fmt.Println(library.SearchByID(1))
  fmt.Println(library.SearchByName("Sword of Destiny"))

  // Замена генератора идентификатора
  library.ChangeId(generateRandomly)
  library.AddBook("Blood of Elves", "Andrzej Sapkowski")
  fmt.Println(library.SearchByName("Blood of Elves"))

  // Map
  mapStorage := NewMapStorage()
  library = CreateLibrary(mapStorage, generateRandomly)
  library.AddBook("Time of Contempt", "Andrzej Sapkowski")
  library.AddBook("Baptism of Fire", "Andrzej Sapkowski")

  fmt.Println(library.SearchByName("Baptism of Fire"))
  fmt.Println(library.SearchByName("Time of Contempt"))
}
