package taskdb

import (
	"log"

	"github.com/boltdb/bolt"
)

// TODO: rewrite to use the Task struct

// Task Holds a todo list task
type Task struct {
	Desc []byte
	ID   int
	done bool
}

// Init Create a task database
func Init() (*TaskDB, error) {
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
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		e := b.Put(id, task)
		return e
	})
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	db, err := Init()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// }
