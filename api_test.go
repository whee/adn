package adn

import (
	"testing"
)

func TestGet(t *testing.T) {
	p, err := GetPost("511604")
	if err != nil {
		t.Error(`GetPost("511604")`, err)
	}

	if p.User.Username != "whee" {
		t.Error("Post username is incorrect")
	}
}
