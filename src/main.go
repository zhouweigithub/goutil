package main

import (
	"fmt"
	"time"

	timeutil "github.com/zhouweigithub/goutil/timeUtil"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(1666236746, 0))
	fmt.Println(timeutil.FormateDateTimeString(time.Unix(1666236746, 0)))
}
