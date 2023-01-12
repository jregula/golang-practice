package db

import (
	"encoding/binary"

	bolt "github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err

	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		// Persist bytes to users bucket.
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return 0, nil

}

func DeleteTask(taskNumber []int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		for _, task := range taskNumber {
			b.Delete(itob(task))
		}
		return nil
	})
	
	if err != nil {
		return err
	}
	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(taskBucket).Cursor()
	
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}