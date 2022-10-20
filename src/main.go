package main

import (
	"fmt"
	timeutil "goutil/timeUtil"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(1666236746, 0))
	fmt.Println(timeutil.FormateDateTimeString(time.Unix(1666236746, 0)))
}
