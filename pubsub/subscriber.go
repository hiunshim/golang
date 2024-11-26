package main

import "fmt"

// UserRegisteredHandler handles the user registered event
func UserRegisteredHandler(eventChan <-chan Event) {
	for event := range eventChan {
		userRegisteredEvent, ok := event.Data.(UserRegisteredEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		// Handle the event
		fmt.Println("New user registered:")
		fmt.Println("ID:", userRegisteredEvent.ID)
		fmt.Println("Name:", userRegisteredEvent.Name)
		fmt.Println("Email:", userRegisteredEvent.Email)
	}
}

func UserUpdatedHandler(eventChan <-chan Event) {
	for event := range eventChan {
		userUpdatedEvent, ok := event.Data.(UserUpdatedEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		fmt.Println("User updated:")
		fmt.Println("ID:", userUpdatedEvent.ID)
		fmt.Println("Name:", userUpdatedEvent.Name)
		fmt.Println("Email:", userUpdatedEvent.Email)
	}
}

func UserDeletedHandler(eventChan <-chan Event) {
	for event := range eventChan {
		userDeletedEvent, ok := event.Data.(UserDeletedEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		// Handle the event
		fmt.Println("User deleted:")
		fmt.Println("ID:", userDeletedEvent.ID)
	}
}
