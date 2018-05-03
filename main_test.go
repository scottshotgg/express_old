package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Args = append(os.Args, "program.expr")

	os.Exit(m.Run())
}
