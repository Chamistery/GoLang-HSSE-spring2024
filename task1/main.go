package main

import (
	"fmt"
	"math/rand"
	generator "task1/generator"
	library "task1/library"
	storage "task1/storage"
)

func IncrementalIDGenerator(lastID int) int {
	return lastID + 1
}

func RandomIDGenerator(lastID int) int {
	return rand.Intn(100000)
}

func main() {
	// Slice
	sliceStorage := storage.NewSliceStorage()
	idGenerator := generator.NewIDGenerator(IncrementalIDGenerator)
	libr := library.CreateLibrary(sliceStorage, idGenerator)

	libr.AddBook("The Last Wish", "Andrzej Sapkowski")
	libr.AddBook("Sword of Destiny", "Andrzej Sapkowski")

	fmt.Println(libr.SearchByID(1))
	fmt.Println(libr.SearchByName("Sword of Destiny"))

	// Замена генератора идентификатора на случайный
	idGenerator = generator.NewIDGenerator(RandomIDGenerator)
	libr.ChangeId(idGenerator)
	libr.AddBook("Blood of Elves", "Andrzej Sapkowski")
	fmt.Println(libr.SearchByName("Blood of Elves"))
	fmt.Println(libr.SearchByName("Sword of Destiny"))

	// Map
	mapStorage := storage.NewMapStorage()
	libr = library.CreateLibrary(mapStorage, idGenerator)

	libr.AddBook("Time of Contempt", "Andrzej Sapkowski")
	libr.AddBook("Baptism of Fire", "Andrzej Sapkowski")

	fmt.Println(libr.SearchByName("Baptism of Fire"))
	fmt.Println(libr.SearchByName("Time of Contempt"))
}
