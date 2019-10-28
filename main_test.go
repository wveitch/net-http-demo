package main

import (
	"flag"
	"testing"
)

var systemTest *bool
func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
}

func TestM(t *testing.T)  {
	if *systemTest {
		main()
	}
}
