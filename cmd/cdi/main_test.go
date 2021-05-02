package main_test

import (
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	c := testcli.Command("cdi")
	c.Run()

	if !c.Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
	}

	if !c.StdoutContains("haz \"123\" branches") {
		t.Fatalf("Expected %q to contain %q", c.Stdout(), "haz \"123\" branches")
	}
}

func TestSomething(t *testing.T) {

	assert.True(t, true, "True is true!")

}
