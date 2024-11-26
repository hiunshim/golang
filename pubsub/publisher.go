package main

import (
	"time"
)

// User represents a user entity
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRegistrationService represents the service responsible for user registration
type UserRegistrationService struct {
	eventBus *EventBus
}

// NewUserRegistrationService creates a new instance of the user registration service
func NewUserRegistrationService(eventBus *EventBus) *UserRegistrationService {
	return &UserRegistrationService{
		eventBus: eventBus,
	}
}

// RegisterUser registers a new user and publishes a user registered event
func (urs *UserRegistrationService) RegisterUser(id int, name, email string) {
	// Simulate user registration
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	// Create the user registered event
	event := Event{
		Type:      "UserRegistered",
		Timestamp: time.Now(),
		Data: UserRegisteredEvent{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	// Publish the event
	urs.eventBus.Publish(event)
}

func (urs *UserRegistrationService) UpdateUser(id int, name, email string) {
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	event := Event{
		Type:      "UserUpdated",
		Timestamp: time.Now(),
		Data: UserUpdatedEvent{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	urs.eventBus.Publish(event)
}

func (urs *UserRegistrationService) DeleteUser(id int) {
	user := User{
		ID: id,
	}
	event := Event{
		Type:      "UserDeleted",
		Timestamp: time.Now(),
		Data: UserDeletedEvent{
			ID: user.ID,
		},
	}
	urs.eventBus.Publish(event)
}
