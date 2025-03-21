package main

import (
	"fmt"

	storage "github.com/jacobmontes14/montyd/internal/datastore"

	"strconv"
)

func main() {
	dataStore := storage.NewDataStore()
	dataStore.AddKeyValue(1, "First")
	dataStore.AddKeyValue(2, "Second")
	dataStore.AddKeyValue(3, "Third")

	fmt.Println(dataStore.GetValue(1))
	fmt.Println(dataStore.GetValue(2))
	fmt.Println(dataStore.GetValue(3))

	slice := dataStore.GetAllKeys()

	for val := range slice {
		fmt.Print(strconv.Itoa(val) + " ")
	}
}
