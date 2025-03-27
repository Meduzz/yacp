package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
)

type Storage struct {
	db *badger.DB
}

var (
	server *Storage
)

func InitStorage(dir string) error {
	opts := badger.DefaultOptions(dir)

	if dir == "" {
		opts = opts.WithInMemory(true)
	}

	db, err := badger.Open(opts)

	if err != nil {
		return err
	}

	server = &Storage{db}

	return nil
}

func SaveChat(chat *Chat) error {
	// serialize payload to json
	bs, err := json.Marshal(chat)

	if err != nil {
		return err
	}

	// store in badger under chats.{chat.id}
	return server.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(fmt.Sprintf("chats.%s", chat.ID)), bs)
	})
}

func LoadChat(id string) (*Chat, error) {
	chatt := &Chat{}

	// load from badger under key chats.{id}
	err := server.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("chats.%s", id)))

		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			// deserialize from json
			return json.Unmarshal(val, chatt)
		})
	})

	return chatt, err
}

func RemoveChat(id string) error {
	// remove key chats.{id} from badger
	return server.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(fmt.Sprintf("chats.%s", id)))
	})
}

func ListChats() ([]*Chat, error) {
	chatts := make([]*Chat, 0)

	// scan badger for keys with prefix chats.
	err := server.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.IteratorOptions{
			Prefix: []byte(fmt.Sprintf("chats.")),
		})

		it.Rewind()

		for it.Valid() {
			chatt := &Chat{}
			err := it.Item().Value(func(val []byte) error {
				// desserialize each of them from json
				return json.Unmarshal(val, chatt)
			})

			if err == nil {
				chatts = append(chatts, chatt)
			} else {
				log.Printf("Loading item from iterator threw error: %v\n", err)
			}

			it.Next()
		}

		it.Close()

		return nil
	})

	return chatts, err
}
