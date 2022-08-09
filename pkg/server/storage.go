package server

import lru "github.com/hashicorp/golang-lru"

type Store struct {
	*lru.Cache
}

var store *Store

func GetStore() *Store {
	return store
}

func init() {
	ret, _ := lru.New(128)
	store = &Store{ret}
}
