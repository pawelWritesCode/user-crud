package manager

import (
	"fmt"
	"sync"

	"github.com/pawelWritesCode/user-crud/internal/database"
	"github.com/pawelWritesCode/user-crud/internal/model"
)

var mu sync.Mutex

// UsersManager is service that has ability to work with User entity.
// db field represents persistence layer
type UsersManager struct {
	db *database.DB
}

func NewUsersManager(db *database.DB) UsersManager {
	return UsersManager{db: db}
}

// Create saves provided user into the db
// Method is safe for concurrent usage
func (um UsersManager) Create(u model.User) (model.User, error) {
	for _, user := range um.db.Users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return model.User{}, fmt.Errorf("user exists in database")
		}
	}

	mu.Lock()
	defer mu.Unlock()
	um.db.Index++
	u.Id = um.db.Index

	um.db.Users = append(um.db.Users, u)

	return u, nil
}

// Delete removes user from db matching provided user id.
// Method is safe for concurrent usage
func (um UsersManager) Delete(userId int) error {
	newUsers := make([]model.User, 0, len(um.db.Users))
	mu.Lock()
	defer mu.Unlock()
	for i, user := range um.db.Users {
		if user.Id == userId {

			newUsers = append(um.db.Users[:i], um.db.Users[i+1:]...)
			um.db.Users = newUsers

			return nil
		}
	}

	return fmt.Errorf("could not find in database user of id %+v", userId)
}

// Replace replaces user definition within db matching provided user id with provided user
// Method is safe for concurrent usage
func (um UsersManager) Replace(userId int, to model.User) error {
	newUsers := make([]model.User, 0, len(um.db.Users))
	mu.Lock()
	defer mu.Unlock()
	for i, user := range um.db.Users {
		if user.Id == userId {
			newUsers = append(um.db.Users[:i], um.db.Users[i+1:]...)

			to.Id = userId
			newUsers = append(newUsers, to)
			um.db.Users = newUsers

			return nil
		}
	}

	return fmt.Errorf("could not find in database user of id %+v", userId)
}

// FindOne returns user entity matching provided user id.
// Method is safe for concurrent usage
func (um UsersManager) FindOne(userId int) (model.User, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range um.db.Users {
		if user.Id == userId {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user of id %d does not exist in database", userId)
}

// GetAll returns all users currently persisted in db.
// Method is safe for concurrent usage
func (um UsersManager) GetAll() []model.User {
	mu.Lock()
	defer mu.Unlock()
	return um.db.Users
}
