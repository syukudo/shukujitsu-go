package main

import (
	"fmt"

	"github.com/syukudo/shukujitsu-go"
)

func main() {
	entries, err := shukujitsu.AllEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		fmt.Printf("%s = %s\n", e.YMD, e.Name)
		fmt.Printf("Year = %d\n", e.Year)
		fmt.Printf("Month = %d\n", e.Month)
		fmt.Printf("Day = %d\n", e.Day)
	}
}
