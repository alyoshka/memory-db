package db

import "errors"

type value struct {
	val    int64
	isNull bool
}

// Database - simple in-memory database
type Database struct {
	data    map[string]int64
	history []map[string]value
}

func NewDatabase() *Database {
	return &Database{
		data:    make(map[string]int64),
		history: []map[string]value{},
	}
}

// Begin - creates new transaction
func (db *Database) Begin() {
	db.history = append(db.history, make(map[string]value))
}

// Commit - clears the transaction history
func (db *Database) Commit() {
	db.history = []map[string]value{}
}

// Get - get key's value, if key is not presented in database - return error
func (db *Database) Get(key string) (int64, error) {
	value, ok := db.data[key]
	if ok {
		return value, nil
	}
	return 0, errors.New("NULL")
}

// NumEqualTo - return the number of variables that are currently set to value
func (db *Database) NumEqualTo(value int64) int64 {
	var result int64
	for _, v := range db.data {
		if v == value {
			result++
		}
	}
	return result
}

// Rollback - Undo all of the commands issued in the most recent transaction block and close the block
func (db *Database) Rollback() error {
	historyLength := len(db.history)
	if historyLength > 0 {
		for k, v := range db.history[historyLength-1] {
			if v.isNull {
				delete(db.data, k)
			} else {
				db.data[k] = v.val
			}
		}
		db.history = db.history[:historyLength-1]
		return nil
	} else {
		return errors.New("NO TRANSACTION")
	}
}

// Set - set the variable "key" to the value "val"
func (db *Database) Set(key string, val int64) {
	historyLength := len(db.history)
	if historyLength > 0 {
		currentValue, ok := db.data[key]
		db.history[historyLength-1][key] = value{
			val:    currentValue,
			isNull: !ok,
		}
	}
	db.data[key] = val
}

// Unset - Unset the variable "key", making it just like that variable was never set.
func (db *Database) Unset(key string) {
	val, ok := db.data[key]
	if ok {
		historyLength := len(db.history)
		if historyLength > 0 {
			db.history[historyLength-1][key] = value{
				val:    val,
				isNull: false,
			}
		}
		delete(db.data, key)
	}
}
