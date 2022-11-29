package random_id

import (
	"fmt"

	"github.com/google/uuid"
)

func UUIDGenerator() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)
	fmt.Printf("String: %s\n", id.String())
	fmt.Printf("Time: %v\n", id.Time())
	fmt.Printf("ClockSequence: %v\n", id.ClockSequence())
	fmt.Printf("Version: %s\n", id.Version())
	fmt.Printf("ID: %v\n", id.ID())
	fmt.Printf("urn: %s\n", id.URN())
	fmt.Printf("Variant: %s\n", id.Variant())

	fmt.Println("------------------------ String")

	stringId := uuid.NewString()
	fmt.Printf("String UUID: %s\n", stringId)

	fmt.Println("------------------------ Parse")
	newId, err := uuid.Parse("5ac87246-047c-4a5a-b294-77c2ef5c7a42")
	if err != nil {
		fmt.Printf("uuid is not valid: %s", err)
		return
	}
	fmt.Printf("Given string is valid and version is %v\n", newId.Version())

}
