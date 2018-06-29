package model

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"todos/database"
)

func UpdateTodoList(bucketName string, dataStruct TodoList) error {
	// Handle DB changes
	err := database.DBCon.Update(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte(bucketName))

		// Find next available id for autoincrementing
		id, _ := b.NextSequence()
		dataStruct.ID = int(id)

		// Serialize TodoList instance
		buf, err := json.Marshal(dataStruct)
		if err != nil {
			return err
		}

		// Update TodoList record in DB
		return b.Put(itob(dataStruct.ID), buf)
	})

	return err
}
