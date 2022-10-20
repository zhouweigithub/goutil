package errutil

import (
	"fmt"
	logutil "goutil/logUtil"
	"runtime/debug"
)

func CatchError() {
	if err := recover(); err != nil {
		stack := string(debug.Stack())
		fmt.Printf("catch error ---> %v\n", err)
		logutil.Error(fmt.Sprintf("catch error ---> %v\nstack trace ---> %s", err, stack))
	}
}
