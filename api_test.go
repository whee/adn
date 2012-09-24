package adn

import (
	"testing"
)

func TestGet(t *testing.T) {
	p := &Post{}
	if err := p.Get("511604"); err != nil {
		t.Error(`Post.Get("511604")`, err)
	}

	if p.User.Username != "whee" {
		t.Error("Post username is incorrect")
	}
}
