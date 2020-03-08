package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// Task Holds a todo list task
type Task struct {
	Desc string
	ID   int
	Done bool
}

// Init Create a task database
func Init() (*bolt.DB, error) {
	// create db
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return db, err
	}

	// create bucket if it does not already exist
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		return err
	})

	return db, nil
}

// AddTask Creates a new task in the database
func AddTask(db *bolt.DB, task string) {
	err := db.Update(func(tx *bolt.Tx) error {
		// Get bucket
		b := tx.Bucket([]byte("Tasks"))
		// Create task
		task := Task{Desc: task, Done: false}

		// Generate id from bucket
		id, _ := b.NextSequence()

		// Add to task as int
		task.ID = int(id)

		// Convert task to buffer
		buf, e := json.Marshal(task)
		if e != nil {
			return e
		}

		// Store in db
		e = b.Put(itob(task.ID), buf)
		return e
	})
	if err != nil {
		log.Fatal(err)
	}
}

// ListTasks Lists all incomplete tasks in database
func ListTasks(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		// If bucket is empty, notify user none exist
		if size := b.Sequence(); size == 0 {
			fmt.Println("No tasks created yet!")
			return nil
		}

		b.ForEach(func(k, v []byte) error {
			var task Task
			err := json.Unmarshal(v, &task)
			if !task.Done {
				fmt.Printf("%v %v\n", task.ID, task.Desc)
			}
			return err
		})
		return nil
	})
}

// CompleteTask Mark existing task as done
func CompleteTask(db *bolt.DB, key int) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		// Get existing task, if it exists
		var task Task
		taskBin := b.Get(itob(key))
		if taskBin == nil {
			fmt.Println("Task not found")
			return nil
		}
		json.Unmarshal(taskBin, &task)

		// Mark task as done and update database
		task.Done = true
		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}
		err = b.Put(itob(key), buf)

		fmt.Println("Task completed")
		return err
	})
}

// Convert int to byte array (width 8)
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
