package main

import (
	"fmt"
	"testing"

	configutil "github.com/zhouweigithub/goutil/configUtil"
	"github.com/zhouweigithub/goutil/threadutil"
)

type ConModel struct {
	Name string
	Sex  string
	Age  int
}

func TestConfig(t *testing.T) {
	var no = ConModel{}
	var err = configutil.ToModel(&no)
	t.Log(no, err)
}

func TestThreading(t *testing.T) {
	var sources = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("start")
	fmt.Println(sources)
	threadutil.Threading(sources, 2, func(item *int) {
		//fmt.Println(*item)
		*item = *item + 10
	})
	fmt.Println(sources)
	fmt.Println("over")
}
