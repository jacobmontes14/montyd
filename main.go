package main

import (
	"fmt"

	storage "github.com/jacobmontes14/montyd/internal/datastore"
)

func main() {
	dataStore := storage.NewDataStore()
	dataStore.AddKeyValue(1, "First")
	dataStore.AddKeyValue(2, "Second")
	dataStore.AddKeyValue(3, "Third")

	fmt.Println(dataStore.GetValue(1))
	fmt.Println(dataStore.GetValue(2))
	fmt.Println(dataStore.GetValue(3))
}
