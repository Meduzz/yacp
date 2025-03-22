package storage

import (
	"log"

	"github.com/dgraph-io/badger"
)

type BadgerStorage struct {
	db *badger.DB
}

func NewInMemoryBadgerStorage() *BadgerStorage {
	opts := badger.DefaultOptions("")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open Badger DB: %v", err)
	}
	return &BadgerStorage{db}
}

func NewFileBasedBadgerStorage(file string) *BadgerStorage {
	opts := badger.DefaultOptions(file)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open Badger DB: %v", err)
	}
	return &BadgerStorage{db}
}
