package db

import "home-provider/models"

// Database provides thread-safe access to a database of books.
type Database interface {

	// Close closes the database, freeing up any available resources.
	Close()

	GetUserDao() UserDao
}

// UserDao provides thread-safe access to a database of books.
type UserDao interface {
	CreateUser(u *models.User) error
}
