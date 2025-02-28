package main

import (
	"testing"
)

func TestHashing(t *testing.T) {
	password := "securepassword"
	hashed1 := Hashing(password)
	hashed2 := Hashing(password)
	if hashed1 == "" || hashed2 == "" {
		t.Error("Hashing function returned an empty string")
	}
	if hashed1 == password || hashed2 == password {
		t.Error("Hashing function returned plain text password")
	}
	if hashed1 == hashed2 {
		t.Error("Hashing function produced identical hashes for the same password (should be different due to salting)")
	}
}
