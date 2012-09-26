// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

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
