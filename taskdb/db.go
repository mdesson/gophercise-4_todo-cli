package taskdb

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Create db if it does not exist
	db, err := bolt.Open("taskdb.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
