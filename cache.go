package cache

import (
	"time"
)

type Pair struct {
	key      string
	value    string
	deadline time.Time
}

type Cache struct {
	Dict []Pair
}

func NewCache() Cache {
	return Cache{}
}

func (cash *Cache) Get(key string) (string, bool) {
	for _, pair := range cash.Dict {
		if pair.key == key && pair.deadline.IsZero() {
			return pair.value, true
		} else if pair.key == key && time.Now().Before(pair.deadline) {
			return pair.value, true
		}
	}
	return "", false
}

func (cash *Cache) Put(key, value string) {
	for i, pair := range cash.Dict {
		if pair.key == key {

			cash.Dict[i].value = value
			return
		}
	}
	newPair := Pair{key, value, time.Time{}}
	cash.Dict = append(cash.Dict, newPair)
}

func (cash *Cache) Keys() []string {
	var keys []string
	for _, pair := range cash.Dict {
		if pair.deadline.IsZero() {
			keys = append(keys, pair.key)
		} else if time.Now().Before(pair.deadline) {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (cash *Cache) PutTill(key, value string, deadline time.Time) {
	newPair := Pair{key, value, deadline}
	cash.Dict = append(cash.Dict, newPair)
}
