package random_id

import (
	"fmt"

	"github.com/jaevor/go-nanoid"
)

func NanoIDGenerator() {
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	id1 := canonicID()
	fmt.Printf("ID 1 with length 21 chars: %s\n", id1)

	decenaryID, err := nanoid.CustomASCII("0123456789ab", 13)
	if err != nil {
		panic(err)
	}

	id2 := decenaryID()
	fmt.Printf("ID 2: %s\n", id2)
}
