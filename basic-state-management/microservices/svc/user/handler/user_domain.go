package handler

// vo
type UserID string
type UserName string

// entity
type User struct {
	ID   UserID
	Name UserName
}

// aggregate root
type AggregateRootUser = User
