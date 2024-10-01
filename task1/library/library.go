package library

import (
	generator "task1/generator"
	storage "task1/storage"
)

type Library struct {
	storage storage.BookStorage
	idGen   *generator.IDGenerator
}

func CreateLibrary(storage storage.BookStorage, idGen *generator.IDGenerator) *Library {
	storage.SetIDGenerator(idGen)
	return &Library{
		storage: storage,
	}
}

func (lib *Library) ChangeId(idGen *generator.IDGenerator) {
	lib.idGen = idGen
	lib.storage.RefreshID(idGen)
}

func (lib *Library) AddBook(name, author string) {
	lib.storage.AddBook(name, author)
}

func (lib *Library) SearchByID(id int) *storage.Book {
	return lib.storage.SearchByID(id)
}

func (lib *Library) SearchByName(name string) *storage.Book {
	return lib.storage.SearchByName(name)
}
