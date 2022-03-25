package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	fmt.Printf("The time %d %d", delay, text)
}
