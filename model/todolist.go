package model

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"todos/database"
	"todos/util"
)

func TodoListUpdate(bucketName string, dataStruct TodoList) error {
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
		return b.Put(util.Itob(dataStruct.ID), buf)
	})

	return err
}

func GetTodoList(bucketName string) ([]TodoList, error) {

	var todolist_array []TodoList

	err := database.DBCon.View(func(tx *bolt.Tx) error {
		// Get TodoBucket instance
		b := tx.Bucket([]byte("TodoBucket"))
		c := b.Cursor()

		// Get all data in TodoBucket
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var todolist TodoList
			// Unseriale each row in TodoBucket
			if err := json.Unmarshal(v, &todolist); err != nil {
				panic(err)
			}
			todolist_array = append(todolist_array, todolist)
		}

		return nil
	})

	return todolist_array, err
}

func RemoveTodoList(bucketName string, ID int) error {
	// Handle DB changes
	err := database.DBCon.Update(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte(bucketName))

		// Delete Key from bucket
		return b.Delete(util.Itob(ID))
	})

	return err
}
