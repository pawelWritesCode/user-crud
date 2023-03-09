package repository

import "github.com/pawelWritesCode/user-crud/internal/model"

// User represents set of method available on User entity
type User interface {
	// Create creates a new User
	Create(u model.User) (model.User, error)

	// Delete removes a User by its user ID
	Delete(userId int) error

	// Replace allows to replace User entity with provided
	Replace(userId int, to model.User) error

	// FindOne returns User matching provided user ID
	FindOne(userId int) (model.User, error)

	// GetAll returns all users
	GetAll() []model.User
}
