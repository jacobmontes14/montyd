package storage

type KeyValue struct {
	Key   int
	Value string
}

type Storage struct {
	keyStore map[int]string
}

func NewDataStore() *Storage {
	s := Storage{keyStore: make(map[int]string)}
	return &s
}

func (storage *Storage) GetValue(key int) string {
	return storage.keyStore[key]
}

func (storage *Storage) AddKeyValue(key int, value string) {
	storage.keyStore[key] = value
}

func (storage *Storage) RemoveKey(key int) {
	delete(storage.keyStore, key)
}

func (storage *Storage) GetAllKeys() []KeyValue {
	toReturn := []KeyValue{}
	for key, val := range storage.keyStore {
		toReturn = append(toReturn, KeyValue{Key: key, Value: val})
	}

	return toReturn
}
