package main

import (
	"fmt"
	"testing"

	configutil "github.com/zhouweigithub/goutil/configUtil"
	"github.com/zhouweigithub/goutil/queryutil"
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

type Model struct {
	Name string
	Age  int
}

var sources = []Model{}

func init() {
	sources = append(sources, Model{Name: "liming1", Age: 10})
	sources = append(sources, Model{Name: "liming2", Age: 11})
	sources = append(sources, Model{Name: "liming3", Age: 12})
	sources = append(sources, Model{Name: "liming4", Age: 13})
	sources = append(sources, Model{Name: "liming5", Age: 11})
	sources = append(sources, Model{Name: "liming6", Age: 15})
	sources = append(sources, Model{Name: "liming7", Age: 16})
}

func TestFilter(t *testing.T) {
	// var a = queryutil.First(sources, func(item *Model) bool { return item.Age > 13 })
	// a.Name = "hello world"
	// fmt.Println(*a)
	// fmt.Println(sources[4])
	// var b = queryutil.Last(sources, func(item *Model) bool { return item.Name == "liming5" })
	// fmt.Println(*b)
	// var c = queryutil.Contains(sources, func(item *Model) bool { return item.Age == 18 })
	// fmt.Println(c)
	// var d = queryutil.Where(sources, func(item *Model) bool { return item.Age < 15 })
	// d[0].Name = "hello world"
	// fmt.Println(*d[0])
	// fmt.Println(sources[0])
	// var e = queryutil.Select(sources, func(item *Model) string { return item.Name })
	// fmt.Println(e)

	// var x = queryutil.Distinct(sources, func(item *Model) int { return item.Age })
	// fmt.Println(x)

	queryutil.OrderBy(sources, func(i, j int) bool { return sources[i].Age < sources[j].Age })
}
