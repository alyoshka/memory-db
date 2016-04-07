package db

import (
	"testing"
)

func TestEmptyVal(t *testing.T) {
	database := NewDatabase()
	_, err := database.Get("a")
	if err == nil {
		t.Error("Get error")
	}
}

func TestGet(t *testing.T) {
	database := NewDatabase()
	database.Set("a", 42)
	result, err := database.Get("a")
	if err != nil {
		t.Error("Get error")
	}
	if result != 42 {
		t.Error("Invalid result")
	}
}

func TestNumEqualTo(t *testing.T) {
	database := NewDatabase()
	result := database.NumEqualTo(42)
	if result != 0 {
		t.Error("Invalid result")
	}
	database.Set("a", 42)
	result = database.NumEqualTo(42)
	if result != 1 {
		t.Error("Invalid result")
	}
	database.Begin()
	database.Set("b", 42)
	result = database.NumEqualTo(42)
	if result != 2 {
		t.Error("Invalid result")
	}
	database.Rollback()
	result = database.NumEqualTo(42)
	if result != 1 {
		t.Error("Rollback error")
	}
}

func TestSetRollback(t *testing.T) {
	database := NewDatabase()
	database.Set("a", 42)
	result, err := database.Get("a")
	if err != nil {
		t.Error("Get error")
	}
	if result != 42 {
		t.Error("Invalid result")
	}
	database.Begin()
	database.Set("a", 43)
	database.Set("b", 55)
	result, err = database.Get("a")
	if err != nil {
		t.Error("Get error")
	}
	if result != 43 {
		t.Error("Invalid result")
	}
	result, err = database.Get("b")
	if err != nil {
		t.Error("Get error")
	}
	if result != 55 {
		t.Error("Invalid result")
	}
	database.Rollback()
	result, err = database.Get("a")
	if err != nil {
		t.Error("Get error")
	}
	if result != 42 {
		t.Error("Invalid result")
	}
	result, err = database.Get("b")
	if err == nil {
		t.Error("Set rollback error")
	}
}

func TestUnset(t *testing.T) {
	database := NewDatabase()
	database.Set("a", 42)
	result, err := database.Get("a")
	if err != nil {
		t.Error("Get error")
	}
	if result != 42 {
		t.Error("Invalid result")
	}
	database.Unset("a")
	result, err = database.Get("a")
	if err == nil {
		t.Error("Unset error")
	}
}

func TestRollbackUnset(t *testing.T) {
	database := NewDatabase()
	database.Set("a", 42)
	database.Begin()
	database.Unset("a")
	_, err := database.Get("a")
	if err == nil {
		t.Error("Unset error")
	}
	database.Rollback()
	result, err := database.Get("a")
	if err != nil || result != 42 {
		t.Error("Rollback error")
	}
}

func TestNoTransaction(t *testing.T) {
	database := NewDatabase()
	err := database.Rollback()
	if err == nil {
		t.Error("No transaction")
	}
}

func TestCommit(t *testing.T) {
	database := NewDatabase()
	database.Set("a", 42)
	database.Begin()
	database.Set("a", 43)
	database.Commit()
	result, err := database.Get("a")
	if err != nil || result != 43 {
		t.Error("Commit error")
	}
}
