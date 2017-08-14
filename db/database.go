package db

// Database provides thread-safe access to a database of books.
type Database interface {

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}
