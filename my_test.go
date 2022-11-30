package main

import (
	"testing"

	configutil "github.com/zhouweigithub/goutil/configUtil"
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
