package taskdb

import (
	"encoding/binary"
	"encoding/json"
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

// TODO: List all (incomplete) tasks

// TODO: Complete task, given a key

// Convert int to byte array (width 8)
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// func main() {
// 	db, err := Init()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	AddTask(db, "do the thing1")
// }
