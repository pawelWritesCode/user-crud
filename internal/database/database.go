package database

import "github.com/pawelWritesCode/user-crud/internal/model"

// DB represents minimalistic database
type DB struct {
	Index int
	Users []model.User
}

func New() *DB {
	return &DB{Users: []model.User{}, Index: 0}
}
