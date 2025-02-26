package testing

import (
	"WorkoutTracker/api"
	"github.com/alexedwards/argon2id"
	"testing"
)

func TestHashing(t *testing.T) {
	input := "mypassword"
	hashedPassword := main.hashing("mypassword")
	match, _ := argon2id.ComparePasswordAndHash(input, hashedPassword)
	if !match {
		t.Error("Invalid hash")
	}
}

func TestSession(t *testing.T) {
	t.Skip("Skipping for now")
}
