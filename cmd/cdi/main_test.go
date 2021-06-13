package main_test

import (
	"github.com/rendon/testcli"
	"os"
	"testing"
)

func TestGitHub_Branch_count(t *testing.T) {

	os.Setenv("TESTING", "true")
	defer os.Unsetenv("TESTING")

	repo := "https://github.com/random-org/amazing-repo.git"
	c := testcli.Command("cdi", repo)

	c.Run()

	if !c.Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
	}

	if !c.StdoutContains("haz 23 branches") {
		t.Fatalf("Expected %q to contain %q", c.Stdout(), "haz 23 branches\n")
	}
}
