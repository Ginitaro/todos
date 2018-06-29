package database

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var DBCon *bolt.DB

const Bucket = "TodoBucket"

func Connect() {

	// Open the DB connection
	var err error
	DBCon, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize TodoBucket bucket
	DBCon.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(Bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_ = b
		return nil
	})

	// Print out existing key/value for debugging
	DBCon.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(Bucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
