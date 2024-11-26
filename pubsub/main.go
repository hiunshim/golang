package main

import (
	"time"
)

func main() {
	eventBus := NewEventBus()

	userRegisteredChan := make(chan Event)
	userUpdatedChan := make(chan Event)
	userDeletedChan := make(chan Event)

	eventBus.Subscribe("UserRegistered", userRegisteredChan)
	eventBus.Subscribe("UserUpdated", userUpdatedChan)
	eventBus.Subscribe("UserDeleted", userDeletedChan)

	go UserRegisteredHandler(userRegisteredChan)
	go UserUpdatedHandler(userUpdatedChan)
	go UserDeletedHandler(userDeletedChan)

	userService := NewUserRegistrationService(eventBus)

	userService.RegisterUser(1, "John Doe", "john.doe@example.com")
	userService.UpdateUser(1, "John Smith", "john.smith@example.com")
	userService.DeleteUser(1)

	time.Sleep(time.Second)
}
