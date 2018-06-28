package database

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
)

var DBCon *bolt.DB

//TODO: FIX dataStruct when its passed
func UpdateTodoData(bucketName string, dataStruct TodoData) error {
	// Handle DB changes
	err := DBCon.Update(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte(bucketName))

		// Find next available id for autoincrementing
		id, _ := b.NextSequence()
		dataStruct.ID = int(id)

		// Serialize TodoData instance
		buf, err := json.Marshal(dataStruct)
		if err != nil {
			return err
		}

		// Update TodoData record in DB
		return b.Put(itob(dataStruct.ID), buf)
	})

	return err
}

func UpdateTodo(bucketName string, TodoDataID int, dataStruct Todo) error {
	// Handle DB changes
	err := DBCon.Update(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte(bucketName))

		// Get TodoData JSON by id from TodoBucket
		v := b.Get([]byte(itob(TodoDataID)))

		var todo_data_item TodoData

		// Unserialize JSON to Tododata instance
		if err := json.Unmarshal(v, &todo_data_item); err != nil {
			panic(err)
		}

		// Find next available id for autoincrementing
		if len(todo_data_item.Todos) > 0 {
			dataStruct.ID = todo_data_item.Todos[len(todo_data_item.Todos)-1].ID + 1
		} else {
			dataStruct.ID = 0
		}

		// Append Todo into Todos of the TodoData instance
		todo_data_item.Todos = append(todo_data_item.Todos, dataStruct)

		// Serialize TodoData instance
		buf, err := json.Marshal(todo_data_item)
		if err != nil {
			return err
		}

		// Update TodoData record in DB
		return b.Put(itob(todo_data_item.ID), buf)
	})

	return err
}

type Todo struct {
	ID   int
	Name string
	Done bool
}

type TodoData struct {
	ID    int
	Title string
	Todos []Todo
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
