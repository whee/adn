// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

import (
	"strings"
)

type Scopes []string

const (
	ScopeBasic     = "basic"
	ScopeStream    = "stream"
	ScopeEmail     = "email"
	ScopeWritePost = "write_post"
	ScopeFollow    = "follow"
	ScopeMessages  = "messages"
	ScopeExport    = "export"
)

func (s Scopes) Spaced() string {
	return strings.Join(s, " ")
}

func (s Scopes) String() string {
	return strings.Join(s, ",")
}
