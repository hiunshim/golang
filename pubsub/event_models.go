package main

type UserRegisteredEvent struct {
	ID    int
	Name  string
	Email string
}

type UserUpdatedEvent struct {
	ID    int
	Name  string
	Email string
}

type UserDeletedEvent struct {
	ID int
}
