package random_id

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func OtherUUIDGenerator() {

	uuidv4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uuidv4)

	uuid2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		return
	}
	fmt.Printf("Successfully parsed: %s\n", uuid2)
}
