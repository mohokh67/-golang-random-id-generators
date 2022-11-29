package random_id

import (
	"fmt"
	"time"

	// "time"

	"github.com/segmentio/ksuid"
)

func KSUIDGenerator() {
	ksuidExample := ksuid.New()
	fmt.Printf("KSUID example: %s\n", ksuidExample)
	fmt.Printf("KSUID string: %s\n", ksuidExample.String())
	fmt.Println(ksuidExample.Payload())
	fmt.Printf("KSUID time: %s\n", ksuidExample.Time())
	fmt.Printf("KSUID timestamp: %d\n", ksuidExample.Timestamp())

	fmt.Println("------------------------ Next and Prev")

	fmt.Printf("Next ID after KSUID: %s\n", ksuidExample.Next())
	fmt.Printf("Previous ID after KSUID: %s\n", ksuidExample.Prev())

	fmt.Println("------------------------ Compare")

	ksuidYesterday, err := ksuid.NewRandomWithTime(time.Now().AddDate(0, 0, -1))
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	ksuidToday := ksuid.New()
	ksuidTomorrow, _ := ksuid.NewRandomWithTime(time.Now().AddDate(0, 0, +1))
	fmt.Printf("KSUID yesterday: %s\n", ksuidYesterday)
	fmt.Printf("KSUID today: %s\n", ksuidToday)
	fmt.Printf("KSUID tomorrow: %s\n", ksuidTomorrow)

	fmt.Printf("Compare KSUID yesterday and today: %d\n", ksuid.Compare(ksuidYesterday, ksuidToday))

	fmt.Printf("Compare KSUID today and today: %d\n", ksuid.Compare(ksuidToday, ksuidToday))

	fmt.Printf("Compare KSUID today and yesterday: %d\n", ksuid.Compare(ksuidToday, ksuidYesterday))

	fmt.Println("------------------------ Sort")

	ksuidList := []ksuid.KSUID{ksuidTomorrow, ksuidYesterday, ksuidToday}
	listSorted := ksuid.IsSorted(ksuidList)
	fmt.Printf("is list sorted: %t\n", listSorted)

	if !listSorted {
		fmt.Println("lets sort the list...")
		ksuid.Sort(ksuidList)
		fmt.Printf("is list sorted: %t\n", ksuid.IsSorted(ksuidList))
	}

	for _, id := range ksuidList {
		fmt.Println(id.Time())
	}

	fmt.Println("------------------------ Parse")
	ksuidFromString, err := ksuid.Parse("2IBRD17gWUmoY0wgs4Fcvo9SiT2")
	if err != nil {
		fmt.Printf("ksuid is not valid: %s", err)
		return
	}
	fmt.Printf("Given string is valid and KSUID: %s\n", ksuidFromString)
	fmt.Printf("Given string is valid and KSUID time is: %s\n", ksuidFromString.Time())

}
