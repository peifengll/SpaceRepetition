package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, now.Location())
	fmt.Printf("%+v\n", today)

}
